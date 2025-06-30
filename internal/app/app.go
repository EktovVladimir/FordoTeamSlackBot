package app

import (
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/grpc/user_finder"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/handlers/auth_handler"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/handlers/management_handler"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/handlers/review_handler"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/handlers/slack_handler"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/jobs"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/services/gh_service"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/services/jira_service"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/services/review_manager"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/services/slack_service"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/services/user_finder"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/environment"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository"
	"github.com/andygrunwald/go-jira"
	"github.com/google/go-github/v72/github"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	"github.com/uptrace/bun"
	"net/http"
	"sync"
	"time"
)

type app struct {
	cfg     *config.Config
	appName string
	pg      *bun.DB
	redis   *redis.Client
}

func New(cfg *config.Config, appName string) *app {
	return &app{cfg: cfg, appName: appName}
}

func (app *app) Run(ctx context.Context) *sync.WaitGroup {
	wg := &sync.WaitGroup{}

	var err error

	app.pg, err = app.connectPostgres(ctx)
	if err != nil {
		logrus.Errorf("Error connecting to Postgres: %v", err)
		return wg
	}

	app.redis, err = app.connectRedis(ctx)
	if err != nil {
		logrus.Errorf("Error connecting to Redis: %v", err)
		return wg
	}

	userRepo := repository.NewPostgresUserRepository(app.pg)
	settingRepo := repository.NewPostgresSettingRepository(app.pg)
	deploymentRepo := repository.NewPostgresDeploymentRepository(app.pg)
	codeReviewRepo := repository.NewPostgresCodeReviewRepository(app.pg)

	auditJob := jobs.NewAuditLogger(
		app.redis,
		jobs.WithInterval(app.cfg.AuditLogger.Interval),
		jobs.WithExpiry(app.cfg.AuditLogger.Expiry),
		jobs.WithTrim(app.cfg.AuditLogger.TrimCount),
		jobs.WithRetriever[*db.User](userRepo, "users"),
		jobs.WithRetriever[*db.Setting](settingRepo, "settings"),
		jobs.WithRetriever[*db.Deployment](deploymentRepo, "deployments"),
		jobs.WithRetriever[*db.CodeReview](codeReviewRepo, "code-reviews"))

	auditJob.Start(ctx)

	slackClient := slack.New(
		app.cfg.Slack.Token,
		slack.OptionDebug(environment.IsDev),
		slack.OptionHTTPClient(&http.Client{
			Timeout: time.Second * 30,
		}))
	githubClient := github.NewClient(nil).WithAuthToken(app.cfg.Github.Token)

	tp := jira.BasicAuthTransport{
		Username: app.cfg.Jira.Email,
		Password: app.cfg.Jira.Token,
	}
	jiraClient, _ := jira.NewClient(tp.Client(), app.cfg.Jira.BaseUrl)

	userFinder := user_finder.New(userRepo, slackClient, githubClient.Search)
	commitRetriever := gh_service.NewCommitRetriever(githubClient.PullRequests)
	prService := gh_service.NewPullRequestService(githubClient.PullRequests)
	jrService := jira_service.NewIssueService(jiraClient.Issue, app.cfg.Jira.BaseUrl)
	messagePoster := slack_service.NewMessagePoster(slackClient)

	crManager := review_manager.NewReviewManager(codeReviewRepo, userFinder, commitRetriever, prService, jrService, messagePoster)

	managementHandler := management_handler.New(userRepo, settingRepo, deploymentRepo, codeReviewRepo)
	authHandler := auth_handler.New(app.cfg.Auth)
	reviewHandler := review_handler.New(crManager)
	slackHandler := slack_handler.New(crManager, slackClient)

	srv := newServer(
		app,
		managementHandler,
		authHandler,
		reviewHandler,
		slackHandler)

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

	go func() {
		<-ctx.Done()
		logrus.Info("Shutting down application...")

		if err := app.pg.Close(); err != nil {
			logrus.Errorf("Error closing Postgres connection: %v", err)
		}
		if err := app.redis.Close(); err != nil {
			logrus.Errorf("Error closing Redis connection: %v", err)
		}

		logrus.Info("Resources released")
	}()

	return wg
}
