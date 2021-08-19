package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database struct {
		Host     string `json:"host" envconfig:"DATABASE_HOST"`
		Port     string `json:"port" envconfig:"DATABASE_PORT"`
		Username string `json:"username" envconfig:"DATABASE_USERNAME"`
		Password string `json:"password" envconfig:"DATABASE_PASSWORD"`
		Name     string `json:"name" envconfig:"DATABASE_NAME"`
	} `json:"database"`
}

func Init() (*Config, error) {
	var c Config

	if err := envconfig.Process("go-auth", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
