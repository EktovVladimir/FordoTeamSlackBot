package run

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/abstract"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/repository"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/service/deploy"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/service/github"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/service/jira"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/service/settings"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/service/slack"
)

type AppContext struct {
	Infrastructure models.Infrastructure
	DeployService  abstract.DeployService
}

func NewApp() *AppContext {
	ghConfig := getGitHubConfig()
	slackConfig := getSlackConfig()
	jiraConfig := getJiraConfig()

	//БД и репозитории
	//TODO инициализация БД?
	dbRepository := repository.NewDummyRepository()

	//Сервисы инфраструктуры
	infra := models.Infrastructure{
		PrService:       github.New(ghConfig),
		SlackService:    slack.New(slackConfig),
		IssueService:    jira.New(jiraConfig),
		SettingsService: settings.New(dbRepository),
	}

	deployService := deploy.New(infra)

	return &AppContext{
		Infrastructure: infra,
		DeployService:  deployService,
	}
}

func getGitHubConfig() config.Github {
	return config.Github{
		Token: "default", //TODO секреты
		Owner: "EktovVladimir",
	}
}

func getSlackConfig() config.Slack {
	return config.Slack{
		BotToken: "default", //TODO секреты
	}
}

func getJiraConfig() config.Jira {
	return config.Jira{
		Token: "default",             //TODO секреты
		Email: "default@default.com", //TODO секреты
	}
}
