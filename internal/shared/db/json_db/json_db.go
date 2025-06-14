package json_db

import (
	"context"
	"encoding/json"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/db"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"sync"
	"time"
)

const (
	usersFileName       = "users.json"
	settingsFileName    = "settings.json"
	deploymentsFileName = "deployments.json"
	codeReviewsFileName = "code_reviews.json"
)

type store struct {
	cfg         config.JsonStoreConfig
	Users       *db.Entity[*db.User]
	Settings    *db.Entity[*db.Setting]
	Deployments *db.Entity[*db.Deployment]
	CodeReviews *db.Entity[*db.CodeReview]
	mu          sync.Mutex
}

func NewStore(cfg config.JsonStoreConfig) *store {
	return &store{
		cfg:         cfg,
		Users:       db.NewEntity[*db.User](),
		Settings:    db.NewEntity[*db.Setting](),
		Deployments: db.NewEntity[*db.Deployment](),
		CodeReviews: db.NewEntity[*db.CodeReview](),
	}
}

func (s *store) GetUsers() *db.Entity[*db.User] {
	return s.Users
}

func (s *store) GetSettings() *db.Entity[*db.Setting] {
	return s.Settings
}

func (s *store) GetDeployments() *db.Entity[*db.Deployment] {
	return s.Deployments
}

func (s *store) GetCodeReviews() *db.Entity[*db.CodeReview] {
	return s.CodeReviews
}

func (s *store) LoadFromFile() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	path := s.cfg.Path

	users, err := readFormFile[*db.User](path, usersFileName)
	if err != nil {
		return err
	}
	s.Users = db.NewLoadedEntity(users)

	settings, err := readFormFile[*db.Setting](path, settingsFileName)
	if err != nil {
		return err
	}
	s.Settings = db.NewLoadedEntity(settings)

	deployments, err := readFormFile[*db.Deployment](path, deploymentsFileName)
	if err != nil {
		return err
	}
	s.Deployments = db.NewLoadedEntity(deployments)

	codeReviews, err := readFormFile[*db.CodeReview](path, codeReviewsFileName)
	if err != nil {
		return err
	}
	s.CodeReviews = db.NewLoadedEntity(codeReviews)

	return nil
}

func (s *store) SyncFiles() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	path := s.cfg.Path

	if err := writeToFile[*db.User](path, usersFileName, s.Users.GetAll()); err != nil {
		return err
	}

	if err := writeToFile[*db.Setting](path, settingsFileName, s.Settings.GetAll()); err != nil {
		return err
	}

	if err := writeToFile[*db.Deployment](path, deploymentsFileName, s.Deployments.GetAll()); err != nil {
		return err
	}

	if err := writeToFile[*db.CodeReview](path, codeReviewsFileName, s.CodeReviews.GetAll()); err != nil {
		return err
	}

	return nil
}

func (s *store) StartFileSync(ctx context.Context, interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			logrus.Debug("Starting file sync cycle")

			select {
			case <-ticker.C:
				func() {
					defer func() {
						if err := recover(); err != nil {
							logrus.Errorf("Panic during file sync: %v", err)
						}
					}()

					if err := s.SyncFiles(); err != nil {
						logrus.Errorf("Error syncing files: %v", err)
					} else {
						logrus.Debug("File sync completed successfully")
					}
				}()
			case <-ctx.Done():
				logrus.Infof("Stopping file sync. Reason: %v", ctx.Err())
				return
			}
		}
	}()
}

func readFormFile[T any](storePath string, fileName string) ([]T, error) {
	filePath := path.Join(storePath, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.WriteFile(filePath, []byte("[]"), 0644); err != nil {
			return nil, errors.Wrapf(err, "Error creating file %s", filePath)
		}
		return make([]T, 0), nil
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, errors.Wrapf(err, "Error reading file %s", filePath)
	}

	var res []T
	if err := json.Unmarshal(data, &res); err != nil {
		return nil, errors.Wrapf(err, "Error unmarshaling %s", filePath)
	}

	return res, nil
}

func writeToFile[T any](storePath string, fileName string, data []T) error {
	filePath := path.Join(storePath, fileName)

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return errors.Wrapf(err, "Error marshalling %s", fileName)
	}

	return os.WriteFile(filePath, bytes, 0644)
}
