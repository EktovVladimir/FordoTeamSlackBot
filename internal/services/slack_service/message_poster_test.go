package slack_service

import (
	"flag"
	mocks "github.com/EktovVladimir/FordoTeamSlackBot/internal/services/slack_service/mocks"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var update = flag.Bool("update", false, "update golden files")

func TestCreateReviewThread_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	src := mocks.NewMocksource(ctrl)
	src.EXPECT().
		GetChannelId().
		Return("C0000000000")

	b, err := os.ReadFile("testdata/cr_thread_text.golden")
	expectText := normalizeSlackMessage(string(b))

	slClient := mocks.NewMockslackClient(ctrl)
	slClient.EXPECT().
		SendMessageContext(gomock.Any(), "C0000000000", gomock.Any()).
		Return("C0000000000", "1750801202.479019", expectText, nil)

	service := NewMessagePoster(slClient)

	actual, err := service.CreateReviewThread(t.Context(), src, &models.CrTreadMessage{
		Requester: &models.User{
			SlackName: "testUser1",
			Email:     "testUser1@test.com",
		},
		Reviewers: []*models.User{
			{
				SlackName: "testUser2",
				Email:     "testUser2@test.com",
			},
			{
				SlackName: "testUser3",
				Email:     "testUser3@test.com",
			},
		},
		Prs: []*models.PullRequestRef{
			models.NewPullRequestRef("TestOwner", "test-repo", 100),
			models.NewPullRequestRef("TestOwner", "test-repo2", 200),
		},
		Issues: []*models.JiraIssue{
			{
				BaseUrl: "https://test.atlassian.net",
				Title:   "Test Issue text",
				Key:     "PROJ-1",
				Project: "PROJ",
			},
		},
	})

	expected := &models.CreatedMessage{
		ChannelId:    "C0000000000",
		Text:         expectText,
		ResponseText: expectText,
		Ts:           "1750801202.479019",
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func normalizeSlackMessage(msg string) string {
	msg = strings.TrimSpace(msg)
	msg = strings.ReplaceAll(msg, "\r\n", "\n")
	return msg
}
