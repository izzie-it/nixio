package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// Load loads configuration from environment variables to cfg struct
func Load(cfg interface{}) error {
	godotenv.Load()

	return env.Parse(cfg)
}
