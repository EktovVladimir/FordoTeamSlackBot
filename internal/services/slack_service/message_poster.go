//go:generate mockgen -destination=mocks/mock.go -source message_poster.go

package slack_service

import (
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/utils/slackutils"
	"github.com/slack-go/slack"
	"strings"
)

type source interface {
	GetChannelId() string
}

type slackClient interface {
	SendMessageContext(context.Context, string, ...slack.MsgOption) (string, string, string, error)
	DeleteMessageContext(context.Context, string, string) (string, string, error)
}

type MessagePoster struct {
	slClient slackClient

	_client *slack.Client
}

func NewMessagePoster(slClient slackClient) *MessagePoster {
	res := &MessagePoster{
		slClient: slClient,
	}

	if client, ok := slClient.(*slack.Client); ok {
		res._client = client
	}

	return res
}

func (m *MessagePoster) CreateReviewThread(ctx context.Context, src source, message *models.CrTreadMessage) (*models.CreatedMessage, error) {
	sb := strings.Builder{}

	sb.WriteString("#cr ")
	sb.WriteString(slackutils.UserMentions(message.Reviewers...))
	sb.WriteString(slackutils.N())

	if len(message.Issues) > 1 {
		sb.WriteString(slackutils.B("Tasks: "))
	} else {
		sb.WriteString(slackutils.B("Task: "))
	}

	sb.WriteString(slackutils.JiraIssueList(message.Issues...))
	sb.WriteString(slackutils.N())

	if len(message.Prs) > 1 {
		sb.WriteString(slackutils.B("PRs: "))
	} else {
		sb.WriteString(slackutils.B("PR: "))
	}

	sb.WriteString(slackutils.PullRequestList(message.Prs...))

	text := sb.String()

	meta := slack.SlackMetadata{
		EventType: "code_review_requested",
		EventPayload: map[string]any{
			"prs": message.Prs,
		},
	}

	channelId, ts, respText, err := m.slClient.SendMessageContext(
		ctx,
		src.GetChannelId(),
		slack.MsgOptionPost(),
		slack.MsgOptionText(text, false),
		slack.MsgOptionMetadata(meta))
	if err != nil {
		return nil, err
	}

	res := &models.CreatedMessage{
		Ts:           ts,
		ChannelId:    channelId,
		Text:         text,
		ResponseText: respText,
		Meta:         meta.EventPayload,
	}

	return res, nil
}

func (m *MessagePoster) DeleteMessage(ctx context.Context, src source, ts string) error {
	_, _, err := m.slClient.DeleteMessageContext(ctx, src.GetChannelId(), ts)
	return err
}
