package management_handler

import (
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/handlers/management_handler/code_review_handler"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/handlers/management_handler/deployment_handler"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/handlers/management_handler/settings_handler"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/handlers/management_handler/user_handler"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository"
)

type contextKey string

const apiContextKey contextKey = "app"

type Handler struct {
	cfg                config.ServerConfig
	userRepo           repository.UserRepository
	userHandler        *user_handler.Handler
	settingHandler     *settings_handler.Handler
	deploymentsHandler *deployment_handler.Handler
	codeReviewsHandler *code_review_handler.Handler
}

func New(
	cfg config.ServerConfig,
	userRepo repository.UserRepository,
	settingRepo repository.SettingsRepository,
	deploymentRepo repository.DeploymentRepository,
	codeReviewRepo repository.CodeReviewRepository) *Handler {

	return &Handler{
		cfg:                cfg,
		userRepo:           userRepo,
		userHandler:        user_handler.New(userRepo),
		settingHandler:     settings_handler.New(settingRepo),
		deploymentsHandler: deployment_handler.New(deploymentRepo),
		codeReviewsHandler: code_review_handler.New(codeReviewRepo),
	}
}

type requestContext struct {
	api *Handler
}

func (h *Handler) contextWithApi(ctx context.Context) context.Context {
	return context.WithValue(ctx, apiContextKey, &requestContext{h})
}
