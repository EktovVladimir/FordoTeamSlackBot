package slackutils

import (
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models"
	"strings"
)

func UserMentions(users ...*models.User) string {
	list := make([]string, 0)
	for _, user := range users {
		m := UserMention(user)
		if m != "" {
			list = append(list, m)
		}
	}

	return strings.Join(list, ", ")
}

func UserMention(user *models.User) string {
	text := ""
	if user.SlackName != "" {
		text = Mention(user.SlackName)
	} else if user.SlackId != "" {
		text = Mention(user.SlackId)
	} else if user.GithubLogin != "" {
		text = user.GithubLogin
	}
	return text
}

func B(text string) string {
	return fmt.Sprintf("*%s*", text)
}

func N() string {
	return "\n"
}

func Link(url, text string) string {
	return fmt.Sprintf("<%s|%s>", url, text)
}

func Mention(userName string) string {
	return fmt.Sprintf("<@%s>", userName)
}

func SubteamMention(teamId string) string {
	return fmt.Sprintf("<!subteam^%s>", teamId)
}

func OrdList(list ...string) string {
	sb := &strings.Builder{}
	for i, v := range list {
		sb.WriteString(fmt.Sprintf("%s%d. %s", N(), i+1, v))
	}

	return sb.String()
}

func InlineList(list ...string) string {
	return strings.Join(list, ", ")
}

func JiraIssueList(issues ...*models.JiraIssue) string {
	urls := make([]string, 0)
	for _, v := range issues {
		urls = append(urls, Link(v.GetUrl(), v.Title))
	}

	return OrdList(urls...)
}

func PullRequestList(prs ...*models.PullRequestRef) string {
	urls := make([]string, 0)
	for _, v := range prs {
		urls = append(urls, Link(v.GetUrl(), v.Repo))
	}

	return InlineList(urls...)
}
