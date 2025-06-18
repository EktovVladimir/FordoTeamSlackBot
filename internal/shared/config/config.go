package config

import (
	"context"
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/environment"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"github.com/heetch/confita/backend/flags"
	"path/filepath"
)

type Config struct {
	Log       LogConfig       `config:"log"`
	Server    ServerConfig    `config:"server"`
	JsonStore JsonStoreConfig `config:"jsonStore"`
}

type LogConfig struct {
	Level  string `config:"log-level"`
	Dir    string `config:"log-dir"`
	MaxAge int    `config:"log-maxAge"`
}

type ServerConfig struct {
	Host            string `config:"server-host"`
	Port            int    `config:"server-port"`
	ReadTimeoutSec  int    `config:"server-readTimeoutSec"`
	WriteTimeoutSec int    `config:"server-writeTimeoutSec"`
}

type JsonStoreConfig struct {
	Path string `config:"jsonStore-path"`
}

func Load(appName string) *Config {
	configPath := filepath.Join("configs", fmt.Sprintf("%s.%s.json", appName, environment.Env))

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
		Server: ServerConfig{
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
