package review_handler

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/services/review_manager"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/helpers/api_helper"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/api"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	reviewManager *review_manager.ReviewManager
}

func New(reviewManager *review_manager.ReviewManager) *Handler {
	return &Handler{reviewManager}
}

func (h *Handler) SetupRoutes(router *gin.RouterGroup) {
	router.POST("/create", h.requestReview)
}

func (h *Handler) requestReview(c *gin.Context) {
	var req RequestReviewRequest
	if err := api_helper.ValidateRequest(c, &req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.NewErrorResponse(err.Error()))
		return
	}

	prRefs, err := utils.ParsePullRequestRefFromUrlMany(req.PrUrls)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.NewErrorResponse(err.Error()))
		return
	}

	src := &apiSource{
		channelId: req.ChannelId,
		email:     req.Email,
	}

	cr, err := h.reviewManager.RequestReview(c.Request.Context(), src, prRefs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}

	resp := api.Response[RequestReviewResponse]{
		Data: RequestReviewResponse{
			ThreadTs: cr.Message.Ts,
		},
	}

	c.JSON(http.StatusOK, resp)
}
