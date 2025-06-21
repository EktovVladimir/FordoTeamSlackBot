package auth_handler

import (
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/helpers/api_helper"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/api"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/auth"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	cfg config.AuthConfig
}

func New(cfg config.AuthConfig) *Handler {
	return &Handler{cfg: cfg}
}

func (h *Handler) SetupRoutes(router *gin.RouterGroup) {
	router.POST("/token", h.token)
}

func (h *Handler) token(c *gin.Context) {
	var req TokenRequest
	if err := api_helper.ValidateRequest(c, &req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, api.NewErrorResponse(err.Error()))
		return
	}

	if !h.validateUser(req) {
		c.JSON(http.StatusUnauthorized, api.NewErrorResponse("Invalid username or password"))
		return
	}

	jwt := utils.NewJwtUtils(&utils.JwtConfig{
		Secret: h.cfg.Secret,
		Expiry: h.cfg.Expiry,
	})

	token, err := jwt.GenerateToken(auth.Payload{
		Username: req.Username,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, api.NewErrorResponse(err.Error()))
		return
	}

	resp := api.Response[TokenResponse]{
		Data: TokenResponse{
			Token: token,
		},
	}

	c.JSON(http.StatusOK, resp)

}

func (h *Handler) validateUser(req TokenRequest) bool {
	return req.Username == h.cfg.UserName && req.Password == h.cfg.Password
}
