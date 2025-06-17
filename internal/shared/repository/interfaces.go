package repository

import (
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/types"
)

type UserRepository interface {
	GetAll(context.Context) ([]*db.User, error)
	GetById(context.Context, types.UniqId) (*db.User, error)
	Create(context.Context, *db.User) error
	Update(context.Context, *db.User) error
	Delete(context.Context, types.UniqId) error
}

type SettingsRepository interface {
	GetAll(context.Context) ([]*db.Setting, error)
	GetById(context.Context, types.UniqId) (*db.Setting, error)
	GetByKey(context.Context, string) (*db.Setting, error)
	Create(context.Context, *db.Setting) error
	Update(context.Context, *db.Setting) error
	Delete(context.Context, types.UniqId) error
	DeleteByKey(context.Context, string) error
}

type DeploymentRepository interface {
	GetAll(context.Context) ([]*db.Deployment, error)
	GetById(context.Context, types.UniqId) (*db.Deployment, error)
	Create(context.Context, *db.Deployment) error
	Update(context.Context, *db.Deployment) error
	Delete(context.Context, types.UniqId) error
}

type CodeReviewRepository interface {
	GetAll(context.Context) ([]*db.CodeReview, error)
	GetById(context.Context, types.UniqId) (*db.CodeReview, error)
	Create(context.Context, *db.CodeReview) error
	Update(context.Context, *db.CodeReview) error
	Delete(context.Context, types.UniqId) error
}
