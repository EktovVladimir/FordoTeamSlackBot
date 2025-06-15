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
	Log           LogConfig           `config:"log"`
	ManagementApi ManagementApiConfig `config:"managementApi"`
	JsonStore     JsonStoreConfig     `config:"jsonStore"`
}

type LogConfig struct {
	Level  string `config:"log-level"`
	Dir    string `config:"log-dir"`
	MaxAge int    `config:"log-maxAge"`
}

type ManagementApiConfig struct {
	Host            string `config:"managementApi-host"`
	Port            int    `config:"managementApi-port"`
	ReadTimeoutSec  int    `config:"managementApi-readTimeoutSec"`
	WriteTimeoutSec int    `config:"managementApi-writeTimeoutSec"`
}

type JsonStoreConfig struct {
	Path string `config:"jsonStore-path"`
}

func Load() *Config {
	configPath := filepath.Join("configs", "app."+environment.Env+".json")

	loader := confita.NewLoader(
		file.NewBackend(configPath),
		flags.NewBackend(),
		env.NewBackend())

	cfg := getDefaultConfig()

	if err := loader.Load(context.Background(), cfg); err != nil {
		//Очень не ожидаем получить тут ошибку, но если получили, то совсем всё плохо
		panic(err)
	}

	return cfg
}

func getDefaultConfig() *Config {
	return &Config{
		Log: LogConfig{
			Level:  "info",
			Dir:    "./logs",
			MaxAge: 90,
		},
		ManagementApi: ManagementApiConfig{
			Host:            "localhost",
			Port:            8000,
			ReadTimeoutSec:  60,
			WriteTimeoutSec: 60,
		},
		JsonStore: JsonStoreConfig{
			Path: "./data",
		},
	}
}
