package app

import (
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/grpc/user_finder"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/handlers/auth_handler"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/handlers/management_handler"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/services/user_finder"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/db_adapter/json_db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/environment"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository"
	"github.com/google/go-github/v72/github"
	"github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
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

	userRepo := repository.NewJsonUserRepository(store)
	settingRepo := repository.NewJsonSettingRepository(store)
	deploymentRepo := repository.NewJsonDeploymentRepository(store)
	codeReviewRepo := repository.NewJsonCodeReviewRepository(store)

	slackClient := slack.New(app.cfg.Slack.Token, slack.OptionDebug(environment.IsDev))
	githubClient := github.NewClient(nil).WithAuthToken(app.cfg.Github.Token)

	userFinder := user_finder.New(userRepo, slackClient, githubClient.Search)

	managementHandler := management_handler.New(userRepo, settingRepo, deploymentRepo, codeReviewRepo)
	authHandler := auth_handler.New(app.cfg.Auth)

	srv := newServer(
		app,
		managementHandler,
		authHandler)

	wg.Add(1)
	go func() {
		defer wg.Done()
		srv.Start(ctx)
	}()

	userFinderGrpc := user_finder_server.New(userFinder)

	gSrv := newGrpcServer(
		app,
		userFinderGrpc,
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		gSrv.Start(ctx)
	}()

	return wg
}
