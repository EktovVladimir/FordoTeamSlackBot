package app

import (
	"context"
	"github.com/redis/go-redis/v9"
)

func (app *app) connectRedis(ctx context.Context) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     app.cfg.Redis.Connection,
		Password: app.cfg.Redis.Password,
		DB:       app.cfg.Redis.DataBase,
	})

	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
