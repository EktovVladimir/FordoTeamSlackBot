package config

import (
	"fmt"
	"github.com/EktovVladimir/FordoTeamSlackBot/internal/shared/environment"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Config struct {
	Auth       AuthConfig       `mapstructure:"auth"`
	Log        LogConfig        `mapstructure:"log"`
	Server     ServerConfig     `mapstructure:"server"`
	GrpcServer GrpcServerConfig `mapstructure:"grpcServer"`
	JsonStore  JsonStoreConfig  `mapstructure:"jsonStore"`
	Slack      SlackConfig      `mapstructure:"slack"`
	Github     GithubConfig     `mapstructure:"github"`
	Jira       JiraConfig       `mapstructure:"jira"`
	Mongo      MongoDb          `mapstructure:"mongo"`
}

type AuthConfig struct {
	Secret   string        `mapstructure:"secret"`
	Expiry   time.Duration `mapstructure:"expiry"`
	UserName string        `mapstructure:"userName"`
	Password string        `mapstructure:"password"`
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

type GrpcServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type JsonStoreConfig struct {
	Path string `mapstructure:"path"`
}

type SlackConfig struct {
	Token string `mapstructure:"token"`
}

type GithubConfig struct {
	Owner string `mapstructure:"owner"`
	Token string `mapstructure:"token"`
}

type JiraConfig struct {
	BaseUrl string `mapstructure:"baseUrl"`
	Email   string `mapstructure:"email"`
	Token   string `mapstructure:"token"`
}

type MongoDb struct {
	Connection   string `mapstructure:"connection"`
	MainDatabase string `mapstructure:"mainDbName"`
}

func Load(appName string) *Config {
	configPath := filepath.Join("configs", fmt.Sprintf("%s.%s.json", appName, environment.Env))

	setDefaults()

	viper.SetConfigFile(configPath)
	viper.SetConfigType("json")
	viper.SetEnvPrefix(appName)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if !os.IsNotExist(err) {
			panic(fmt.Errorf("fatal error reading config file: %w", err))
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("fatal error unmarshaling config: %w", err))
	}

	return &cfg
}

func setDefaults() {
	viper.SetDefault("log.level", "debug")
	viper.SetDefault("log.dir", "./logs")
	viper.SetDefault("log.maxAge", 90)

	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.readTimeout", "60s")
	viper.SetDefault("server.writeTimeout", "60s")

	viper.SetDefault("grpcServer.host", "localhost")
	viper.SetDefault("grpcServer.port", 8081)

	viper.SetDefault("jsonStore.path", "./data")
}
