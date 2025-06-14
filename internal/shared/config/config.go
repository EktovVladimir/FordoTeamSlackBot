package config

import (
	"context"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/environment"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
	"path/filepath"
)

type Config struct {
	Log LogConfig `config:"log"`
}

type LogConfig struct {
	Level  string `config:"log-level"`
	Dir    string `config:"log-dir"`
	MaxAge int    `config:"log-max_age"`
}

func Load(ctx context.Context) (*Config, error) {
	configPath := filepath.Join("configs", "app."+environment.Env+".yaml")

	loader := confita.NewLoader(
		file.NewBackend(configPath),
		flags.NewBackend(),
		env.NewBackend())

	cfg := &Config{
		Log: LogConfig{
			Level:  "info",
			Dir:    "./logs",
			MaxAge: 90,
		},
	}

	if err := loader.Load(ctx, cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
