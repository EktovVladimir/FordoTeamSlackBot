package code_review_handler

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/helpers/api_helper"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/api"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	repo repository.CodeReviewRepository
}

func New(repo repository.CodeReviewRepository) *Handler {
	return &Handler{repo}
}

func (h *Handler) SetupRoutes(router *gin.RouterGroup) {
	group := router.Group("/reviews")
	group.GET("", h.getCodeReviews)
	group.GET("/{id}", h.getCodeReview)
	group.POST("", h.createCodeReview)
	group.PUT("/{id}", h.updateCodeReview)
	group.DELETE("/{id}", h.deleteCodeReview)
}

func (h *Handler) getCodeReviews(c *gin.Context) {
	data, err := h.repo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}

	var respData []CodeReviewResponse
	if err := api_helper.MapAndResponse(c, &data, &respData); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}
}

func (h *Handler) getCodeReview(c *gin.Context) {
	id, err := api_helper.GetUniqIdFromRoute(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.NewErrorResponse(err.Error()))
		return
	}

	data, err := h.repo.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, api.NewErrorResponse(err.Error()))
		return
	}

	var respData CodeReviewResponse
	if err := api_helper.MapAndResponse(c, &data, &respData); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}
}

func (h *Handler) createCodeReview(c *gin.Context) {
	var req CreateCodeReviewRequest
	var dbModel db.CodeReview
	if err := api_helper.ValidateAndMapRequest(c, &req, &dbModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.NewErrorResponse(err.Error()))
		return
	}

	if err := h.repo.Create(c.Request.Context(), &dbModel); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}

	var respModel CodeReviewResponse
	if err := api_helper.MapAndResponse(c, &dbModel, &respModel); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}
}

func (h *Handler) updateCodeReview(c *gin.Context) {
	id, err := api_helper.GetUniqIdFromRoute(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.NewErrorResponse(err.Error()))
		return
	}

	var req UpdateCodeReviewRequest
	if err = api_helper.ValidateRequest(c, &req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.NewErrorResponse(err.Error()))
		return
	}

	dbModel, err := h.repo.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, api.NewErrorResponse(err.Error()))
		return
	}

	if err = api_helper.MapIgnoreEmpty(&req, dbModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.NewErrorResponse(err.Error()))
		return
	}

	if err = h.repo.Update(c.Request.Context(), dbModel); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}

	var respModel CodeReviewResponse
	if err = api_helper.MapAndResponse(c, &dbModel, &respModel); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}
}

func (h *Handler) deleteCodeReview(c *gin.Context) {
	id, err := api_helper.GetUniqIdFromRoute(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.NewErrorResponse(err.Error()))
		return
	}

	_, err = h.repo.GetById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, api.NewErrorResponse(err.Error()))
		return
	}

	if err = h.repo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}

	c.Status(http.StatusNoContent)
}
