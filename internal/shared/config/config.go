package config

import (
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/environment"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"time"
)

type Config struct {
	Log       LogConfig       `mapstructure:"log"`
	Server    ServerConfig    `mapstructure:"server"`
	JsonStore JsonStoreConfig `mapstructure:"jsonStore"`
}

type LogConfig struct {
	Level  string `mapstructure:"level"`
	Dir    string `mapstructure:"dir"`
	MaxAge int    `mapstructure:"maxAge"`
}

type ServerConfig struct {
	Host         string        `mapstructure:"host"`
	Port         int           `mapstructure:"port"`
	ReadTimeout  time.Duration `mapstructure:"readTimeout"`
	WriteTimeout time.Duration `mapstructure:"writeTimeout"`
}

type JsonStoreConfig struct {
	Path string `mapstructure:"path"`
}

func Load(appName string) *Config {
	configPath := filepath.Join("configs", fmt.Sprintf("%s.%s.json", appName, environment.Env))

	setDefaults()

	viper.SetConfigFile(configPath)
	viper.SetConfigType("json")
	viper.AutomaticEnv()

	viper.ReadInConfig()

	if err := viper.ReadInConfig(); err != nil {
		if !os.IsNotExist(err) {
			panic(fmt.Errorf("fatal error reading config file: %w", err))
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("fatal error unmarshaling config: %w", err))
	}

	fmt.Println(cfg)

	return &cfg
}

func setDefaults() {
	viper.SetDefault("log.level", "debug")
	viper.SetDefault("log.dir", "./logs")
	viper.SetDefault("log.maxAge", 90)

	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", 8000)
	viper.SetDefault("server.readTimeout", "60s")
	viper.SetDefault("server.writeTimeout", "60s")

	viper.SetDefault("jsonStore.path", "./data")
}
