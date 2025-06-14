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
	Log     LogConfig     `config:"log"`
	CrudApi CrudApiConfig `config:"crudApi"`
}

type LogConfig struct {
	Level  string `config:"log-level"`
	Dir    string `config:"log-dir"`
	MaxAge int    `config:"log-maxAge"`
}

type CrudApiConfig struct {
	Host string `config:"crudApi-host"`
	Port int    `config:"crudApi-port"`
}

func Load() (*Config, error) {
	configPath := filepath.Join("configs", "app."+environment.Env+".json")

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
		CrudApi: CrudApiConfig{
			Host: "localhost",
			Port: 8000,
		},
	}

	if err := loader.Load(context.Background(), cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
