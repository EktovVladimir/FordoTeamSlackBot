package jira

import (
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/issues"
)

type issueService struct {
	config config.Jira
	//TODO клиент jira
}

func New(config config.Jira) *issueService {
	//TODO создание клиента Jira
	return &issueService{
		config: config,
	}
}

func (service *issueService) Get(number string) (issues.Issue, error) {
	return issues.Issue{
		Code:  number,
		Title: "example",
		Url:   fmt.Sprintf("https://%s.atlassian.net/browse/%s", service.config.Workspace, number),
	}, nil
}

func (service *issueService) GetList(numbers []string) ([]issues.Issue, error) {
	res1, _ := service.Get("1001")
	res2, _ := service.Get("1002")
	return []issues.Issue{res1, res2}, nil
}
