package management_api

import (
	"context"
	"errors"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/management_api/handlers"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/repository"
)

type contextKey string

const apiContextKey contextKey = "management_api"

type ManagementApi struct {
	cfg                config.ManagementApiConfig
	userRepo           repository.UserRepository
	userHandler        *handlers.UserManagement
	settingHandler     *handlers.SettingsManagement
	deploymentsHandler *handlers.DeploymentsManagement
	codeReviewsHandler *handlers.CodeReviewsManagement
}

func New(
	cfg config.ManagementApiConfig,
	userRepo repository.UserRepository,
	settingRepo repository.SettingsRepository,
	deploymentRepo repository.DeploymentRepository,
	codeReviewRepo repository.CodeReviewRepository) *ManagementApi {

	return &ManagementApi{
		cfg:                cfg,
		userRepo:           userRepo,
		userHandler:        handlers.NewUserManagement(userRepo),
		settingHandler:     handlers.NewSettingsManagement(settingRepo),
		deploymentsHandler: handlers.NewDeploymentsManagement(deploymentRepo),
		codeReviewsHandler: handlers.NewCodeReviewsManagement(codeReviewRepo),
	}
}

type requestContext struct {
	api *ManagementApi
}

func (a *ManagementApi) contextWithApi(ctx context.Context) context.Context {
	return context.WithValue(ctx, apiContextKey, &requestContext{a})
}

func apiFromContext(ctx context.Context) (*requestContext, error) {
	val := ctx.Value(apiContextKey)
	if rc, ok := val.(*requestContext); ok {
		return rc, nil
	}
	return nil, errors.New("ManagementApi not found in context")
}
