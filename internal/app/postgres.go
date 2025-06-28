package app

import (
	"context"
	"database/sql"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/environment"
	dbmodels "github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/models/db"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func (app *app) connectPostgres(ctx context.Context) (*bun.DB, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(app.cfg.Postgres.Connection)))

	db := bun.NewDB(sqldb, pgdialect.New())

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	db.RegisterModel((*dbmodels.CodeReviewToPR)(nil))

	if environment.IsDev {
		db.AddQueryHook(bundebug.NewQueryHook(
			bundebug.WithEnabled(true),
			bundebug.WithVerbose(true),
			bundebug.WithWriter(logrus.StandardLogger().Writer()),
		))
	}

	return db, nil
}
