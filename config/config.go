package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// AppConfig -.
type AppConfig struct {
	Name         string   `env-required:"true" env:"APP_NAME"    yaml:"name"   `
	Version      string   `env-required:"true" env:"APP_VERSION" yaml:"version" `
	AllowOrigins []string `env-required:"true" env:"ALLOW_ORIGINS" yaml:"allowed_origins"`
}

// HttpConfig -.
type HttpConfig struct {
	Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
}

// RdPostgresConfig -.
type RdPostgresConfig struct {
	Host     string `env-required:"true"  yaml:"host" env:"PG_HOST"`
	User     string `env-required:"true" yaml:"user" env:"PG_USER"`
	Password string `env-required:"true" yaml:"password" env:"PG_PASSWORD"`
	DbName   string `env-required:"true" yaml:"dbname" env:"PG_DB_NAME"`
	Port     string `env-required:"true" yaml:"port" env:"PG_PORT"`
	SslMode  string `yaml:"ssl_mode" env:"PG_SSL_MODE"`
	TimeZone string `yaml:"time_zone" env:"PG_TIMEZONE"`
}

// RdRedisConfig -.
type RdRedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
}

// Config -.
type Config struct {
	App        AppConfig        `yaml:"app"`
	Http       HttpConfig       `yaml:"http"`
	RdPostgres RdPostgresConfig `yaml:"postgres"`
	RdRedis    RdRedisConfig    `yaml:"rd_redis"`
}

// EnvEnum 环境变量枚举配置
type EnvEnum struct {
	Dev  string
	Test string
	Prod string
}

var Env = &EnvEnum{
	Dev:  "dev",
	Test: "test",
	Prod: "prod",
}

var (
	RdPostgres *RdPostgresConfig
	RdRedis    *RdRedisConfig
	App        *AppConfig
	Http       *HttpConfig
)

func init() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	switch env {
	case Env.Dev:
		config := &Config{}

		pwd, _ := os.Getwd()
		configFilePath := filepath.Join(pwd, "config", "development.yaml")

		dataBytes, err := os.ReadFile(configFilePath)
		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(dataBytes, config)
		if err != nil {
			panic(err)
		}

		RdPostgres = &config.RdPostgres
		RdRedis = &config.RdRedis
		App = &config.App
		Http = &config.Http
		break

	case Env.Test:
		panic("还未指定测试环境")

	case Env.Prod:
		panic("还未指定生成环境")

	default:
		panic("意料之外的环境")
	}
}
