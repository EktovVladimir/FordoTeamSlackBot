package jira_service

import (
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models"
	"github.com/andygrunwald/go-jira"
)

type jiraIssueService interface {
	GetWithContext(context.Context, string, *jira.GetQueryOptions) (*jira.Issue, *jira.Response, error)
}

type IssueService struct {
	jiraClient jiraIssueService
	baseUrl    string
}

func NewIssueService(jiraClient jiraIssueService, baseUrl string) *IssueService {
	return &IssueService{jiraClient, baseUrl}
}

func (s *IssueService) GetIssueList(ctx context.Context, keys []string) ([]*models.JiraIssue, error) {

	res := make([]*models.JiraIssue, 0)
	for _, key := range keys {
		jIssue, _, err := s.jiraClient.GetWithContext(ctx, key, nil)
		if err != nil {
			return nil, err
		}

		res = append(res, &models.JiraIssue{
			Key:     jIssue.Key,
			Title:   jIssue.Fields.Summary,
			Project: jIssue.Fields.Project.Key,
			BaseUrl: s.baseUrl,
		})
	}
	return res, nil
}
