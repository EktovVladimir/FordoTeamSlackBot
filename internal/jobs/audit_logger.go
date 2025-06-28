package jobs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

const (
	redisKeyPrefix = "audit-Data"
)

type BetweenRetriever[T auditableEntity] interface {
	GetUpdatedBetween(context.Context, time.Time, time.Time) ([]T, error)
}

type auditableEntity interface {
	db.Auditable
	db.HasUniqId
}

type AuditLoggerOption func(*auditLoggerConfig)
type logReturnerFunc func(context.Context, *auditLoggerIterationContext) ([]*auditLogEntry, error)

type auditLoggerIterationContext struct {
	startTime time.Time
	endTime   time.Time
	mu        *sync.Mutex
}

type auditLogEntry struct {
	UniqId    db.UniqId `json:"uniq_id"`
	Name      string    `json:"Name"`
	Label     string    `json:"Label"`
	UpdatedAt time.Time `json:"updated_at"`
	Data      any       `json:"Data"`
}

type auditLoggerConfig struct {
	interval     time.Duration
	expiry       time.Duration
	trimCount    int64
	logReturners []logReturnerFunc
}

type AuditLogger struct {
	redis *redis.Client
	cfg   *auditLoggerConfig
}

func NewAuditLogger(redis *redis.Client, opts ...AuditLoggerOption) *AuditLogger {
	conf := &auditLoggerConfig{
		expiry:       0,
		interval:     5 * time.Minute,
		logReturners: make([]logReturnerFunc, 0),
	}

	for _, opt := range opts {
		opt(conf)
	}

	return &AuditLogger{
		redis: redis,
		cfg:   conf}
}

func (l *AuditLogger) Start(ctx context.Context) {
	go func() {
		logrus.Info("AuditLogger job started, waiting for initial delay...")

		//time.Sleep(l.cfg.interval)

		ticker := time.NewTicker(l.cfg.interval)
		defer ticker.Stop()

		mu := &sync.Mutex{}

		iterCtx := &auditLoggerIterationContext{
			startTime: time.Now(),
			endTime:   time.Now(),
			mu:        mu,
		}

		for {
			select {
			case <-ticker.C:
				iterCtx.endTime = time.Now()

				l.runIteration(ctx, iterCtx)

				iterCtx.startTime = iterCtx.endTime
			case <-ctx.Done():
				logrus.Infof("Stopping AuditLogger. Reason: %v", ctx.Err())
				return
			}
		}
	}()
}

func (l *AuditLogger) runIteration(ctx context.Context, iterCtx *auditLoggerIterationContext) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Panic in AuditLogger job: %v", r)
		}
	}()

	iterCtx.mu.Lock()
	defer iterCtx.mu.Unlock()

	logrus.Debug("Run auditLogger iteration")

	for _, fn := range l.cfg.logReturners {
		logs, err := fn(ctx, iterCtx)
		if err != nil {
			logrus.Errorf("Error in logReturners: %v", err)
			continue
		}

		if logs != nil && len(logs) > 0 {
			name := logs[0].Name
			logrus.Debugf("Found auditLogger %s Data count: %d", name, len(logs))

			for _, log := range logs {
				key := getAuditLoggerRedisKey(log)
				bytes, err := json.MarshalIndent(log, "", "  ")
				if err != nil {
					logrus.Errorf("Error marshalling %v", err)
					continue
				}
				value := string(bytes)

				l.redis.LPush(ctx, key, value)

				if l.cfg.expiry != 0 {
					l.redis.Expire(ctx, key, l.cfg.expiry)
				}

				if l.cfg.trimCount > 0 {
					l.redis.LTrim(ctx, key, 0, l.cfg.trimCount)
				}
			}
		}
	}
}

func WithInterval(interval time.Duration) AuditLoggerOption {
	return func(config *auditLoggerConfig) {
		config.interval = interval
	}
}

func WithExpiry(expiry time.Duration) AuditLoggerOption {
	return func(config *auditLoggerConfig) {
		config.expiry = expiry
	}
}

func WithTrim(trimCount int64) AuditLoggerOption {
	return func(config *auditLoggerConfig) {
		config.trimCount = trimCount
	}
}

func WithRetriever[T auditableEntity](ret BetweenRetriever[T], name string) AuditLoggerOption {
	return func(cfg *auditLoggerConfig) {

		fn := func(ctx context.Context, iterCtx *auditLoggerIterationContext) ([]*auditLogEntry, error) {
			return scanAndString(ctx, ret, name, iterCtx)
		}

		cfg.logReturners = append(cfg.logReturners, fn)
	}
}

func scanAndString[T auditableEntity](ctx context.Context, ret BetweenRetriever[T], name string, iterCtx *auditLoggerIterationContext) ([]*auditLogEntry, error) {
	if ret == nil {
		return nil, errors.New("retriever cannot be nil")
	}

	items, err := ret.GetUpdatedBetween(ctx, iterCtx.startTime, iterCtx.endTime)
	if err != nil {
		return nil, err
	}

	res := make([]*auditLogEntry, 0)
	for _, item := range items {
		crDate := item.GetCreatedAt()
		upDate := item.GetUpdatedAt()
		uniqId := item.GetId()

		label := "updated"
		if upDate == crDate {
			label = "created"
		}

		res = append(res, &auditLogEntry{
			UniqId:    uniqId,
			Label:     label,
			Name:      name,
			UpdatedAt: upDate,
			Data:      item,
		})
	}

	return res, nil
}

func getAuditLoggerRedisKey(entry *auditLogEntry) string {
	return fmt.Sprintf("%s:%s:%d", redisKeyPrefix, entry.Name, entry.UniqId)
}
