package deploy

import (
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/abstract"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/helpers"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/git"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/integrations/messenger"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/request_source"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/settings"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/models/shared"
	"regexp"
	"strings"
)

type deploySlackNotification struct {
	prService       abstract.PullRequestService
	slackService    abstract.MessengerService
	issueService    abstract.IssueTrackerService
	settingsService abstract.SettingsService
}

func New(
	prService abstract.PullRequestService,
	slackService abstract.MessengerService,
	issueService abstract.IssueTrackerService,
	settingsService abstract.SettingsService) *deploySlackNotification {
	return &deploySlackNotification{
		prService,
		slackService,
		issueService,
		settingsService,
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
		user, err := s.slackService.GetUser(pr.Requester.Email)
		if err == nil {
			slackUser = user
		}

		channelName = s.settingsService.GetStringItem(source, settings.ChannelIdKey)
	}

	userLink := shared.UserLink{GithubUser: pr.Requester, SlackUser: slackUser}

	thread := &messenger.Thread{
		Text: helpers.GetSlackDeployNotificationMainText(userLink, pr),
		Attachments: []messenger.Attachment{
			helpers.GetSlackDeployNotificationAttachment("opened"),
		},
	}

	err := s.slackService.CreateThread(channelName, thread)

	if err != nil {
		panic(err)
	}

	//TODO сохранение в репу

	commits, err := s.prService.GetCommits(pr.Repo, pr.Number)

	if err != nil {
		panic(err)
	}

	taskPrefix := s.settingsService.GetStringItem(source, settings.IssuePrefixKey)
	issueNumbers := extractIssues(commits, taskPrefix)

	issues, err := s.issueService.GetList(issueNumbers)

	if err != nil {
		panic(err)
	}

	issuesText := helpers.GetSlackIssuesList(issues)

	comment := &messenger.Comment{
		Text: issuesText,
	}

	err = s.slackService.CreateComment(channelName, thread.Ts, comment)

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
