package helpers

import (
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/integrations/git"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/integrations/issues"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/integrations/messenger"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/shared"
	"strings"
)

func GetSlackDeployNotificationMainText(author shared.UserLink, pr git.PullRequest) string {
	userMention := GetSlackMention(author)

	var mentionText string
	if userMention != "" {
		mentionText = fmt.Sprintf("%s создал", userMention)
	} else {
		mentionText = "Создан"
	}

	res := fmt.Sprintf("%s pull request <%s|%s>", mentionText, GetPullRequestUrl(pr), pr.Title)

	//TODO остальной текст

	return res
}

func GetSlackDeployNotificationAttachment(status string) messenger.Attachment {
	//TODO дополнить все ветки
	switch status {
	default:
		return messenger.Attachment{
			Text:  ":loading:",
			Color: "default",
		}
	}
}

func GetSlackIssuesList(issues []issues.Issue) string {
	var sb strings.Builder
	for index, issue := range issues {
		sb.WriteString(fmt.Sprintf("%d. <%s|%s>\n", index+1, issue.Url, issue.Title))
	}
	return sb.String()
}

func GetSlackMention(user shared.UserLink) string {
	if user.SlackUser.Name != "" {
		return fmt.Sprintf("<@%s>", user.SlackUser.Name)
	}

	if user.GithubUser.Login != "" {
		return fmt.Sprintf("%s", user.GithubUser.Login)
	}

	return ""
}
