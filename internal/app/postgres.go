package app

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"golang.org/x/net/context"
)

func (app *app) connectPostgres(ctx context.Context) (*bun.DB, error) {
	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(app.cfg.Postgres.Connection)))

	db := bun.NewDB(sqldb, pgdialect.New())

	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	return db, nil
}
