package run

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/request_source"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/repository"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/service/deploy"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/service/github"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/service/jira"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/service/settings"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/service/slack"
)

func Run() {
	//Конфигурации
	ghConfig := getGitHubConfig()
	slackConfig := getSlackConfig()
	jiraConfig := getJiraConfig()

	//БД и репозитории
	//TODO инициализация БД?
	dbRepository := repository.NewDummyRepository()

	//Сервисы
	ghService := github.New(ghConfig)
	slackService := slack.New(slackConfig)
	jiraService := jira.New(jiraConfig)
	settingsService := settings.New(dbRepository)

	deployNotifService := deploy.New(ghService, slackService, jiraService, settingsService)

	dummyGhRepo, _ := ghService.Get("FordoTeamSlackBot", "1")
	source := request_source.Workflow{
		Repo:   "FordoTeamSlackBot",
		Branch: "develop",
		Team:   "backoffice",
	}

	deployNotifService.CreateThread(dummyGhRepo, source)
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
