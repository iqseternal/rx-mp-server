package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type (
	// Config -.
	Config struct {
		App     app     `yaml:"app"`
		Http    http    `yaml:"http"`
		RdPg    rdPg    `yaml:"rd_pg"`
		RdRedis rdRedis `yaml:"rd_redis"`
	}

	// app -.
	app struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// http -.
	http struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// rdPg -.
	rdPg struct {
		Host     string `env-required:"true"  yaml:"host"             env:"PG_HOST"`
		User     string `env-required:"true" yaml:"user" env:"PG_USER"`
		Password string `env-required:"true" yaml:"password" env:"PG_PASSWORD"`
		DbName   string `env-required:"true" yaml:"dbname" env:"PG_DB_NAME"`
		Port     string `env-required:"true" yaml:"port" env:"PG_PORT"`
		SslMode  string `yaml:"ssl_mode" env:"PG_SSL_MODE"`
		TimeZone string `yaml:"time_zone" env:"PG_TIMEZONE"`
	}

	// rdRedis -
	rdRedis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
	}
)

type EnvEnum struct {
	Dev  string
	Test string
	Prod string
}

var (
	Env = &EnvEnum{
		Dev:  "dev",
		Test: "test",
		Prod: "prod",
	}
)

var (
	RdPg    *rdPg
	RdRedis *rdRedis
	App     *app
	Http    *http
)

func init() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	switch env {
	case Env.Dev:
		config := &Config{}

		dataBytes, err := os.ReadFile("config/development.yaml")
		if err != nil {
			panic(err)
		}

		err = yaml.Unmarshal(dataBytes, config)
		if err != nil {
			panic(err)
		}

		RdPg = &config.RdPg
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
