package settings_handler

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/helpers/api_helper"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/api"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	repo repository.SettingsRepository
}

func New(repo repository.SettingsRepository) *Handler {
	return &Handler{repo}
}

func (h *Handler) SetupRoutes(router *gin.RouterGroup) {
	group := router.Group("/settings")
	group.GET("", h.getSettings)
	group.GET("/:id", h.getSetting)
	group.POST("", h.createSetting)
	group.PUT("/:id", h.updateSetting)
	group.DELETE("/:id", h.deleteSetting)
}

// @Summary	Получить все настройки
// @Tags		settings
// @Security	BearerAuth
// @Produce	json
// @Success	200	{object}	api.Response[[]SettingResponse]
// @Failure	500	{object}	api.Response[any]
// @Router		/management/settings [get]
func (h *Handler) getSettings(c *gin.Context) {
	data, err := h.repo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}

	var respData []SettingResponse
	if err := api_helper.MapAndResponse(c, &data, &respData); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}
}

// @Summary	Получить настройку по ID
// @Tags		settings
// @Security	BearerAuth
// @Produce	json
// @Param		id	path		int	true	"ID настройки"
// @Success	200	{object}	api.Response[SettingResponse]
// @Failure	422	{object}	api.Response[any]	"Ошибка валидации ID"
// @Failure	404	{object}	api.Response[any]	"Настройка не найдена"
// @Failure	500	{object}	api.Response[any]
// @Router		/management/settings/{id} [get]
func (h *Handler) getSetting(c *gin.Context) {
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

	var respData SettingResponse
	if err := api_helper.MapAndResponse(c, &data, &respData); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}
}

// @Summary	Создать новую настройку
// @Tags		settings
// @Security	BearerAuth
// @Accept		json
// @Produce	json
// @Param		request	body		CreateSettingRequest	true	"Данные настройки"
// @Success	200		{object}	api.Response[SettingResponse]
// @Failure	422		{object}	api.Response[any]	"Ошибка валидации"
// @Failure	500		{object}	api.Response[any]
// @Router		/management/settings [post]
func (h *Handler) createSetting(c *gin.Context) {
	var req CreateSettingRequest
	var dbModel db.Setting
	if err := api_helper.ValidateAndMapRequest(c, &req, &dbModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.NewErrorResponse(err.Error()))
		return
	}

	if err := h.repo.Create(c.Request.Context(), &dbModel); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}

	var respModel SettingResponse
	if err := api_helper.MapAndResponse(c, &dbModel, &respModel); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}
}

// @Summary	Обновить настройку
// @Tags		settings
// @Security	BearerAuth
// @Accept		json
// @Produce	json
// @Param		id		path		int						true	"ID настройки"
// @Param		request	body		UpdateSettingRequest	true	"Новые данные настройки"
// @Success	200		{object}	api.Response[SettingResponse]
// @Failure	422		{object}	api.Response[any]	"Ошибка валидации"
// @Failure	404		{object}	api.Response[any]	"Настройка не найдена"
// @Failure	500		{object}	api.Response[any]
// @Router		/management/settings/{id} [put]
func (h *Handler) updateSetting(c *gin.Context) {
	id, err := api_helper.GetUniqIdFromRoute(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.NewErrorResponse(err.Error()))
		return
	}

	var req UpdateSettingRequest
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
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}

	if err = h.repo.Update(c.Request.Context(), dbModel); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}

	var respModel SettingResponse
	if err = api_helper.MapAndResponse(c, &dbModel, &respModel); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}
}

// @Summary	Удалить настройку
// @Tags		settings
// @Security	BearerAuth
// @Produce	json
// @Param		id	path	int	true	"ID настройки"
// @Success	204
// @Failure	422	{object}	api.Response[any]	"Ошибка валидации ID"
// @Failure	404	{object}	api.Response[any]	"Настройка не найдена"
// @Failure	500	{object}	api.Response[any]
// @Router		/management/settings/{id} [delete]
func (h *Handler) deleteSetting(c *gin.Context) {
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
