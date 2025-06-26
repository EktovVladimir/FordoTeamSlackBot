package user_handler

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/helpers/api_helper"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/api"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	repo repository.UserRepository
}

func New(repo repository.UserRepository) *Handler {
	return &Handler{repo}
}

func (h *Handler) SetupRoutes(router *gin.RouterGroup) {
	group := router.Group("/users")
	group.GET("", h.getUsers)
	group.GET("/:id", h.getUser)
	group.POST("/search", h.searchUsers)
	group.POST("", h.createUser)
	group.PUT("/:id", h.updateUser)
	group.DELETE("/:id", h.deleteUser)
}

// @Summary	Возвращает список всех пользователей
// @Tags		users
// @Security	BearerAuth
// @Accept		json
// @Produce	json
// @Success	200	{object}	api.Response[[]UserResponse]
// @Failure	500	{object}	api.Response[any]
// @Router		/management/users [get]
func (h *Handler) getUsers(c *gin.Context) {
	c.Request.Context()
	data, err := h.repo.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}

	var respData []UserResponse
	if err := api_helper.MapAndResponse(c, &data, &respData); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}
}

// @Summary Поиск пользователей по дате изменения
// @Description Возвращает список пользователей, измененных в указанный временной период
// @Tags users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body SearchUsersRequest true "Параметры поиска"
// @Success 200 {object} api.Response[[]UserResponse] "Список найденных пользователей"
// @Failure 400 {object} api.Response[any] "Неверный формат запроса"
// @Failure 422 {object} api.Response[any] "Ошибки валидации параметров"
// @Failure 500 {object} api.Response[any] "Внутренняя ошибка сервера"
// @Router /management/users/search [post]
func (h *Handler) searchUsers(c *gin.Context) {
	var req SearchUsersRequest
	if err := api_helper.ValidateRequest(c, &req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.NewErrorResponse(err.Error()))
		return
	}

	if req.UpdatedAtFrom == nil || req.UpdatedAtTo == nil {
		c.JSON(http.StatusUnprocessableEntity, api.NewErrorResponse("start time or end time is nil"))
		return
	}

	data, err := h.repo.GetUpdatedBetween(c.Request.Context(), *req.UpdatedAtFrom, *req.UpdatedAtTo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}

	var respData []UserResponse
	if err := api_helper.MapAndResponse(c, &data, &respData); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}
}

// @Summary	Возвращает информацию о пользователе
// @Tags		users
// @Security	BearerAuth
// @Accept		json
// @Produce	json
// @Param		id	path		int	true	"ID пользователя"
// @Success	200	{object}	api.Response[UserResponse]
// @Failure	422	{object}	api.Response[any]	"Ошибка в запросе"
// @Failure	404	{object}	api.Response[any]	"Пользователь с таким ID не найден"
// @Failure	500	{object}	api.Response[any]
// @Router		/management/users/{id} [get]
func (h *Handler) getUser(c *gin.Context) {
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

	var respData UserResponse
	if err := api_helper.MapAndResponse(c, &data, &respData); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}
}

// @Summary	Добавляет пользователя в систему
// @Tags		users
// @Security	BearerAuth
// @Accept		json
// @Produce	json
// @Param		message	body		CreateUserRequest	true	"Данные пользователя"
// @Success	200		{object}	api.Response[UserResponse]
// @Failure	422		{object}	api.Response[any]	"Ошибка в запросе"
// @Failure	500		{object}	api.Response[any]
// @Router		/management/users [post]
func (h *Handler) createUser(c *gin.Context) {
	var req CreateUserRequest
	var dbModel db.User
	if err := api_helper.ValidateAndMapRequest(c, &req, &dbModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.NewErrorResponse(err.Error()))
		return
	}

	if err := h.repo.Create(c.Request.Context(), &dbModel); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}

	var respModel UserResponse
	if err := api_helper.MapAndResponse(c, &dbModel, &respModel); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}
}

// @Summary	Изменяет информацию о пользователе
// @Tags		users
// @Security	BearerAuth
// @Accept		json
// @Produce	json
// @Param		id	path		int	true	"ID пользователя"
// @Success	200	{object}	api.Response[UserResponse]
// @Failure	422	{object}	api.Response[any]	"Ошибка в запросе"
// @Failure	404	{object}	api.Response[any]	"Пользователь с таким ID не найден"
// @Failure	500	{object}	api.Response[any]
// @Router		/management/users/{id} [put]
func (h *Handler) updateUser(c *gin.Context) {
	id, err := api_helper.GetUniqIdFromRoute(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.NewErrorResponse(err.Error()))
		return
	}

	var req UpdateUserRequest
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

	var respModel UserResponse
	if err = api_helper.MapAndResponse(c, &dbModel, &respModel); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}
}

// @Summary	Удаляет пользователя
// @Tags		users
// @Security	BearerAuth
// @Accept		json
// @Produce	json
// @Param		id	path	int	true	"ID пользователя"
// @Success	204
// @Failure	422	{object}	api.Response[any]	"Ошибка в запросе"
// @Failure	404	{object}	api.Response[any]	"Пользователь с таким ID не найден"
// @Failure	500	{object}	api.Response[any]
// @Router		/management/users/{id} [delete]
func (h *Handler) deleteUser(c *gin.Context) {
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

	if err := h.repo.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}

	c.Status(http.StatusNoContent)
}
