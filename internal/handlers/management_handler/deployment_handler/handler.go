package deployment_handler

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/helpers/api_helper"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/api"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	repo repository.DeploymentRepository
}

func New(repo repository.DeploymentRepository) *Handler {
	return &Handler{repo}
}

func (h *Handler) SetupRoutes(router *gin.RouterGroup) {
	group := router.Group("/deployments")
	group.GET("", h.getDeployments)
	group.GET("/:id", h.getDeployment)
	group.POST("", h.createDeployment)
	group.PUT("/:id", h.updateDeployment)
	group.DELETE("/:id", h.deleteDeployment)
}

// @Summary	Получить список всех деплоев
// @Tags		deployments
// @Security	BearerAuth
// @Produce	json
// @Success	200	{object}	api.Response[[]DeploymentResponse]
// @Failure	500	{object}	api.Response[any]
// @Router		/management/deployments [get]
func (h *Handler) getDeployments(c *gin.Context) {
	data, err := h.repo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}

	var respData []DeploymentResponse
	if err := api_helper.MapAndResponse(c, &data, &respData); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}
}

// @Summary	Получить информацию о деплое
// @Tags		deployments
// @Security	BearerAuth
// @Produce	json
// @Param		id	path		int	true	"ID деплоя"
// @Success	200	{object}	api.Response[DeploymentResponse]
// @Failure	422	{object}	api.Response[any]	"Ошибка валидации ID"
// @Failure	404	{object}	api.Response[any]	"Деплой не найден"
// @Failure	500	{object}	api.Response[any]
// @Router		/management/deployments/{id} [get]
func (h *Handler) getDeployment(c *gin.Context) {
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

	var respData DeploymentResponse
	if err := api_helper.MapAndResponse(c, &data, &respData); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}
}

// @Summary	Создать новый деплой
// @Tags		deployments
// @Security	BearerAuth
// @Accept		json
// @Produce	json
// @Param		request	body		CreateDeploymentRequest	true	"Данные деплоя"
// @Success	200		{object}	api.Response[DeploymentResponse]
// @Failure	422		{object}	api.Response[any]	"Ошибка валидации"
// @Failure	500		{object}	api.Response[any]
// @Router		/management/deployments [post]
func (h *Handler) createDeployment(c *gin.Context) {
	var req CreateDeploymentRequest
	var dbModel db.Deployment
	if err := api_helper.ValidateAndMapRequest(c, &req, &dbModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.NewErrorResponse(err.Error()))
		return
	}

	if err := h.repo.Create(c.Request.Context(), &dbModel); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}

	var respModel DeploymentResponse
	if err := api_helper.MapAndResponse(c, &dbModel, &respModel); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}
}

// @Summary	Обновить информацию о деплое
// @Tags		deployments
// @Security	BearerAuth
// @Accept		json
// @Produce	json
// @Param		id		path		int						true	"ID деплоя"
// @Param		request	body		UpdateDeploymentRequest	true	"Новые данные деплоя"
// @Success	200		{object}	api.Response[DeploymentResponse]
// @Failure	422		{object}	api.Response[any]	"Ошибка валидации"
// @Failure	404		{object}	api.Response[any]	"Деплой не найден"
// @Failure	500		{object}	api.Response[any]
// @Router		/management/deployments/{id} [put]
func (h *Handler) updateDeployment(c *gin.Context) {
	id, err := api_helper.GetUniqIdFromRoute(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.NewErrorResponse(err.Error()))
		return
	}

	var req UpdateDeploymentRequest
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

	var respModel DeploymentResponse
	if err = api_helper.MapAndResponse(c, &dbModel, &respModel); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}
}

// @Summary	Удалить деплой
// @Tags		deployments
// @Security	BearerAuth
// @Produce	json
// @Param		id	path	int	true	"ID деплоя"
// @Success	204
// @Failure	422	{object}	api.Response[any]	"Ошибка валидации ID"
// @Failure	404	{object}	api.Response[any]	"Деплой не найден"
// @Failure	500	{object}	api.Response[any]
// @Router		/management/deployments/{id} [delete]
func (h *Handler) deleteDeployment(c *gin.Context) {
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
