package json_db

import (
	"context"
	"encoding/json"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/db_adapter"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
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

	dirPerm  = 0755
	filePerm = 0644
)

type store struct {
	cfg         config.JsonStoreConfig
	Users       *db_adapter.Entity[*db.User]
	Settings    *db_adapter.Entity[*db.Setting]
	Deployments *db_adapter.Entity[*db.Deployment]
	CodeReviews *db_adapter.Entity[*db.CodeReview]
	mu          sync.Mutex
}

func NewStore(cfg config.JsonStoreConfig) *store {
	return &store{
		cfg:         cfg,
		Users:       db_adapter.NewEntity[*db.User](),
		Settings:    db_adapter.NewEntity[*db.Setting](),
		Deployments: db_adapter.NewEntity[*db.Deployment](),
		CodeReviews: db_adapter.NewEntity[*db.CodeReview](),
	}
}

func (s *store) GetUsers() *db_adapter.Entity[*db.User] {
	return s.Users
}

func (s *store) GetSettings() *db_adapter.Entity[*db.Setting] {
	return s.Settings
}

func (s *store) GetDeployments() *db_adapter.Entity[*db.Deployment] {
	return s.Deployments
}

func (s *store) GetCodeReviews() *db_adapter.Entity[*db.CodeReview] {
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
	s.Users = db_adapter.NewLoadedEntity(users)

	settings, err := readFormFile[*db.Setting](path, settingsFileName)
	if err != nil {
		return err
	}
	s.Settings = db_adapter.NewLoadedEntity(settings)

	deployments, err := readFormFile[*db.Deployment](path, deploymentsFileName)
	if err != nil {
		return err
	}
	s.Deployments = db_adapter.NewLoadedEntity(deployments)

	codeReviews, err := readFormFile[*db.CodeReview](path, codeReviewsFileName)
	if err != nil {
		return err
	}
	s.CodeReviews = db_adapter.NewLoadedEntity(codeReviews)

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
		logrus.Info("File sync service started, waiting for initial delay...")

		time.Sleep(interval)

		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				s.SyncFilesWithLog()
			case <-ctx.Done():
				logrus.Infof("Stopping file sync. Reason: %v", ctx.Err())
				return
			}
		}
	}()
}

func (s *store) SyncFilesWithLog() {
	logrus.Debug("Starting file sync cycle")

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
}

func (s *store) SaveChanges() error {
	return s.SyncFiles()
}

func readFormFile[T any](storePath string, fileName string) ([]T, error) {
	if err := os.MkdirAll(storePath, dirPerm); err != nil {
		return nil, errors.Wrapf(err, "Error creating directory %s", storePath)
	}

	filePath := path.Join(storePath, fileName)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		if err := os.WriteFile(filePath, []byte("[]"), filePerm); err != nil {
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
	if err := os.MkdirAll(storePath, dirPerm); err != nil {
		return errors.Wrapf(err, "Error creating directory %s", storePath)
	}

	filePath := path.Join(storePath, fileName)

	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return errors.Wrapf(err, "Error marshalling %s", fileName)
	}

	return os.WriteFile(filePath, bytes, filePerm)
}
