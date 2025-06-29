package slack_handler

import (
	"context"
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/services/review_manager"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
	"net/http"
	"time"
)

type Handler struct {
	reviewManager *review_manager.ReviewManager
	slackClient   *slack.Client
}

func New(reviewManager *review_manager.ReviewManager, slackClient *slack.Client) *Handler {
	return &Handler{reviewManager, slackClient}
}

func (h *Handler) SetupRoutes(router *gin.RouterGroup) {
	router.POST("/cr", h.requestReview)
}

func (h *Handler) requestReview(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data"})
		return
	}

	command := c.Request.FormValue("command")
	text := c.Request.FormValue("text")
	userId := c.Request.FormValue("user_id")
	channelId := c.Request.FormValue("channel_id")
	responseURL := c.Request.FormValue("response_url")

	logrus.Debugf("command %v\ntext %v\nuserID %v\nchannelId %v\nresponseURL %v\n", command, text, userId, channelId, responseURL)

	c.String(http.StatusOK, "")

	ctx := context.Background()

	go func() {
		userCtx, userCancel := context.WithTimeout(ctx, 5*time.Second)
		defer userCancel()

		slUser, err := h.slackClient.GetUserInfoContext(userCtx, userId)
		if err != nil {
			msg := fmt.Sprintf("Failed to get user info: %v", err)
			logrus.Error(msg)
			_, _ = h.slackClient.PostEphemeralContext(ctx, channelId, userId, slack.MsgOptionText(msg, false))
			return
		}

		src := &slackSource{
			userId:    userId,
			email:     slUser.Profile.Email,
			channelId: channelId,
		}

		prs, err := utils.ParsePullRequestRefFromUrlTextMany(text)
		if err != nil {
			msg := fmt.Sprintf("Failed to parse PR URLs: %v", err)
			logrus.Error(msg)
			_, _ = h.slackClient.PostEphemeralContext(ctx, channelId, userId, slack.MsgOptionText(msg, false))
			return
		}

		reviewCtx, reviewCancel := context.WithTimeout(ctx, 5*time.Second)
		defer reviewCancel()

		_, err = h.reviewManager.RequestReview(reviewCtx, src, prs)
		if err != nil {
			msg := fmt.Sprintf("Failed to create review: %v", err)
			logrus.Error(msg)
			_, _ = h.slackClient.PostEphemeralContext(ctx, channelId, userId, slack.MsgOptionText(msg, false))
			return
		}
	}()
}
