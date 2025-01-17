package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type (
	// Config -.
	Config struct {
		App  App  `yaml:"app"`
		Http Http `yaml:"http"`
		PG   PG   `yaml:"postgres"`
	}

	// App -.
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}

	// Http -.
	Http struct {
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	}

	// PG -.
	PG struct {
		Host     string `env-required:"true"  yaml:"host"             env:"PG_HOST"`
		User     string `env-required:"true" yaml:"user" env:"PG_USER"`
		Password string `env-required:"true" yaml:"password" env:"PG_PASSWORD"`
		DbName   string `env-required:"true" yaml:"dbname" env:"PG_DB_NAME"`
		Port     string `env-required:"true" yaml:"port" env:"PG_PORT"`
		SslMode  string `yaml:"ssl_mode" env:"PG_SSL_MODE"`
		TimeZone string `yaml:"time_zone" env:"PG_TIMEZONE"`
	}
)

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	dataBytes, err := os.ReadFile("config/development.yaml")
	if err != nil {
		fmt.Println("Error reading config.yaml")
		return nil, err
	}

	err = yaml.Unmarshal(dataBytes, cfg)
	if err != nil {
		fmt.Println("Error parsing config.yaml")
		return nil, err
	}

	return cfg, nil
}
