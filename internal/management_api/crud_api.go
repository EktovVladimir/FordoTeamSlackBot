package management_api

import (
	"context"
	"errors"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/management_api/handlers/user"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/db_adapter"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
)

type contextKey string

const apiContextKey contextKey = "management_api"

type store interface {
	GetUsers() *db_adapter.Entity[*db.User]
	GetSettings() *db_adapter.Entity[*db.Setting]
	GetDeployments() *db_adapter.Entity[*db.Deployment]
	GetCodeReviews() *db_adapter.Entity[*db.CodeReview]
	SaveChanges() error
}

type ManagementApi struct {
	cfg         config.ManagementApiConfig
	db          store
	userHandler *user.Handler
}

func New(cfg config.ManagementApiConfig, db store) *ManagementApi {
	return &ManagementApi{
		cfg:         cfg,
		db:          db,
		userHandler: user.New(db)}
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
