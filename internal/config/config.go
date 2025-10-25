package config

import (
	"log"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort  string `env:"SERVER_PORT" envDefault:"8080"`
	DatabaseURL string `env:"DATABASE_URL"`
}

func Load() (*Config, error) {
	godotenv.Load()

	var cfg Config

	return &cfg, env.Parse(&cfg)
}

func MustLoad() *Config {
	cfg, err := Load()

	if err != nil {
		log.Fatal("config load failed:", err)
	}

	return cfg
}
