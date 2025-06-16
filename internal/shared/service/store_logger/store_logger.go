package store_logger

import (
	"context"
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

type dataStore interface {
	GetUsers() []*db.User
	GetSettings() []*db.Setting
	GetDeployments() []*db.Deployment
	GetCodeReviews() []*db.CodeReview
}

type storeLogger struct {
	store         dataStore
	interval      time.Duration
	lastPositions scanPosition
}

type scanPosition struct {
	users       int
	settings    int
	deployments int
	codeReviews int
}

type scanSlices struct {
	users       []*db.User
	settings    []*db.Setting
	deployments []*db.Deployment
	codeReviews []*db.CodeReview
}

func NewStoreLogger(store dataStore, interval time.Duration) *storeLogger {
	return &storeLogger{
		store:         store,
		interval:      interval,
		lastPositions: scanPosition{},
	}
}

func (sl *storeLogger) StartGo(ctx context.Context) {
	go func() {
		sl.Start(ctx)
	}()
}

func (sl *storeLogger) Start(ctx context.Context) {
	sl.lastPositions = scanPosition{
		users:       len(sl.store.GetUsers()),
		settings:    len(sl.store.GetSettings()),
		deployments: len(sl.store.GetDeployments()),
		codeReviews: len(sl.store.GetCodeReviews()),
	}

	for {
		select {
		case <-ctx.Done():
			logrus.Info("storeLogger done by context")
			return
		case <-time.After(sl.interval):
			positions, slices := sl.scan()

			var sb strings.Builder

			if len(slices.users) > 0 {
				fmt.Fprintf(&sb, "Found new user records: %d\n", len(slices.users))
				for _, item := range slices.users {
					fmt.Fprintln(&sb, item)
				}
			}

			if len(slices.settings) > 0 {
				fmt.Fprintf(&sb, "Found new setting records: %d\n", len(slices.settings))
				for _, item := range slices.settings {
					fmt.Fprintln(&sb, item)
				}
			}

			if len(slices.deployments) > 0 {
				fmt.Fprintf(&sb, "Found new deployment records: %d\n", len(slices.deployments))
				for _, item := range slices.deployments {
					fmt.Fprintln(&sb, item)
				}
			}

			if len(slices.codeReviews) > 0 {
				fmt.Fprintf(&sb, "Found new code-review records: %d\n", len(slices.codeReviews))
				for _, item := range slices.codeReviews {
					fmt.Fprintln(&sb, item)
				}
			}

			sl.lastPositions = positions

			logrus.Info(sb.String())
		}
	}
}

func (sl *storeLogger) scan() (scanPosition, scanSlices) {
	users := sl.store.GetUsers()
	usersLen := len(users)

	resSlices := scanSlices{}
	resPos := scanPosition{}

	if usersLen > sl.lastPositions.users {
		resSlices.users = users[sl.lastPositions.users:]
		resPos.users = usersLen
	}

	settings := sl.store.GetSettings()
	settingsLen := len(settings)

	if settingsLen > sl.lastPositions.settings {
		resSlices.settings = settings[sl.lastPositions.settings:]
		resPos.settings = settingsLen
	}

	deployments := sl.store.GetDeployments()
	deploymentsLen := len(deployments)

	if deploymentsLen > sl.lastPositions.deployments {
		resSlices.deployments = deployments[sl.lastPositions.deployments:]
		resPos.deployments = deploymentsLen
	}

	codeReviews := sl.store.GetCodeReviews()
	codeReviewsLen := len(codeReviews)

	if codeReviewsLen > sl.lastPositions.codeReviews {
		resSlices.codeReviews = codeReviews[sl.lastPositions.codeReviews:]
		resPos.codeReviews = codeReviewsLen
	}

	return resPos, resSlices
}
