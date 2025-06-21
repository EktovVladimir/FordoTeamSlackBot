package app

import (
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/handlers/auth_handler"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/handlers/management_handler"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/db_adapter/json_db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type app struct {
	cfg     *config.Config
	appName string
}

func New(cfg *config.Config, appName string) *app {
	return &app{cfg: cfg, appName: appName}
}

func (app *app) Run(ctx context.Context) *sync.WaitGroup {
	wg := &sync.WaitGroup{}

	store := json_db.NewStore(app.cfg.JsonStore)
	if err := store.LoadFromFile(); err != nil {
		logrus.Errorf("Error loading json data files: %v", err)
		return wg
	}

	defer store.SyncFilesWithLog()

	store.StartFileSync(ctx, 5*time.Minute)

	//TODO UoW
	userRepo := repository.NewJsonUserRepository(store)
	settingRepo := repository.NewJsonSettingRepository(store)
	deploymentRepo := repository.NewJsonDeploymentRepository(store)
	codeReviewRepo := repository.NewJsonCodeReviewRepository(store)

	managementHandler := management_handler.New(userRepo, settingRepo, deploymentRepo, codeReviewRepo)
	authHandler := auth_handler.New(app.cfg.Auth)

	srv := newServer(
		app,
		managementHandler,
		authHandler)

	srv.Start(ctx)

	return wg
}
