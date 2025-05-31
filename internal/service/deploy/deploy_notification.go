package deploy

import (
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/abstract"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/helpers"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/git"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/messenger"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/request_source"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/settings"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/shared"
	"regexp"
	"strings"
)

type deploySlackNotification struct {
	infra models.Infrastructure
}

func New(
	infra models.Infrastructure) *deploySlackNotification {
	return &deploySlackNotification{
		infra,
	}
}

func (s *deploySlackNotification) CreateThread(pr git.PullRequest, source abstract.Source) error {

	slackUser := messenger.User{}
	var channelName string
	switch src := source.(type) {
	case request_source.Slack:
		slackUser = src.User
		channelName = src.Channel
	case request_source.Workflow:
		user, err := s.infra.SlackService.GetUser(pr.Requester.Email)
		if err == nil {
			slackUser = user
		}

		channelName = s.infra.SettingsService.GetStringItem(source, settings.ChannelIdKey)
	}

	userLink := shared.UserLink{GithubUser: pr.Requester, SlackUser: slackUser}

	thread := &messenger.Thread{
		Text: helpers.GetSlackDeployNotificationMainText(userLink, pr),
		Attachments: []messenger.Attachment{
			helpers.GetSlackDeployNotificationAttachment("opened"),
		},
	}

	err := s.infra.SlackService.CreateThread(channelName, thread)

	if err != nil {
		panic(err)
	}

	//TODO сохранение в репу

	commits, err := s.infra.PrService.GetCommits(pr.Repo, pr.Number)

	if err != nil {
		panic(err)
	}

	taskPrefix := s.infra.SettingsService.GetStringItem(source, settings.IssuePrefixKey)
	issueNumbers := extractIssues(commits, taskPrefix)

	issues, err := s.infra.IssueService.GetList(issueNumbers)

	if err != nil {
		panic(err)
	}

	issuesText := helpers.GetSlackIssuesList(issues)

	comment := &messenger.Comment{
		Text: issuesText,
	}

	err = s.infra.SlackService.CreateComment(channelName, thread.Ts, comment)

	if err != nil {
		panic(err)
	}

	return nil
}

func extractIssues(commits []git.Commit, taskPrefix string) []string {
	pattern := fmt.Sprintf(`(?i)\b%s-\d+\b`, taskPrefix)
	re := regexp.MustCompile(pattern)

	set := make(map[string]any)

	for _, commit := range commits {
		match := re.FindString(commit.Message)
		upperMatch := strings.ToUpper(match)
		set[upperMatch] = struct{}{}
	}

	var res []string
	for item := range set {
		res = append(res, item)
	}

	return res
}
