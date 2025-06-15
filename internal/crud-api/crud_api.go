package crud_api

import (
	"context"
	"errors"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/config"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/db_adapter"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
)

type contextKey string

const apiContextKey contextKey = "crud_api"

type store interface {
	GetUsers() *db_adapter.Entity[*db.User]
	GetSettings() *db_adapter.Entity[*db.Setting]
	GetDeployments() *db_adapter.Entity[*db.Deployment]
	GetCodeReviews() *db_adapter.Entity[*db.CodeReview]
}

type CrudApi struct {
	cfg config.CrudApiConfig
	db  store
}

func New(cfg config.CrudApiConfig, db store) *CrudApi {
	return &CrudApi{cfg, db}
}

type requestContext struct {
	api *CrudApi
}

func (a *CrudApi) contextWithApi(ctx context.Context) context.Context {
	return context.WithValue(ctx, apiContextKey, &requestContext{a})
}

func apiFromContext(ctx context.Context) (*requestContext, error) {
	val := ctx.Value(apiContextKey)
	if rc, ok := val.(*requestContext); ok {
		return rc, nil
	}
	return nil, errors.New("CrudApi not found in context")
}
