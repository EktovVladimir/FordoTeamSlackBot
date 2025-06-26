package app

import (
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/grpc/user_finder"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/handlers/auth_handler"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/handlers/management_handler"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/handlers/review_handler"
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
	"github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	"sync"
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

	mongo, err := app.connectMongo(ctx)
	if err != nil {
		logrus.Errorf("Error connecting to Mongo: %v", err)
		return wg
	}
	defer mongo.Disconnect(ctx)

	mongoDb := mongo.Database(app.cfg.Mongo.DataBase)

	redis, err := app.connectRedis(ctx)
	if err != nil {
		logrus.Errorf("Error connecting to Redis: %v", err)
		return wg
	}

	userRepo := repository.NewMongoUserRepository(mongoDb)
	settingRepo := repository.NewMongoSettingRepository(mongoDb)
	deploymentRepo := repository.NewMongoDeploymentRepository(mongoDb)
	codeReviewRepo := repository.NewMongoCodeReviewRepository(mongoDb)

	auditJob := jobs.NewAuditLogger(
		redis,
		jobs.WithInterval(app.cfg.AuditLogger.Interval),
		jobs.WithExpiry(app.cfg.AuditLogger.Expiry),
		jobs.WithTrim(app.cfg.AuditLogger.TrimCount),
		jobs.WithRetriever[*db.User](userRepo, "users"),
		jobs.WithRetriever[*db.Setting](settingRepo, "settings"),
		jobs.WithRetriever[*db.Deployment](deploymentRepo, "deployments"),
		jobs.WithRetriever[*db.CodeReview](codeReviewRepo, "code-reviews"))

	auditJob.Start(ctx)

	slackClient := slack.New(app.cfg.Slack.Token, slack.OptionDebug(environment.IsDev))
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

	srv := newServer(
		app,
		managementHandler,
		authHandler,
		reviewHandler)

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
