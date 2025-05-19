package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	DatabaseConnectionString string `env:"DATABASE_URL"`
	ApiServerHost            string `env:"APISERVER_HOST"`
	ApiServerPort            string `env:"APISERVER_PORT"`
}

func New() (*Config, error) {
	cfg, err := env.ParseAs[Config]()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	return &cfg, nil
}

func (c *Config) GetAddr() string {
	return fmt.Sprintf("%s:%s", c.ApiServerHost, c.ApiServerPort)
}
