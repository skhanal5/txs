package config

import (
	"fmt"
	"os"
)

type Config struct {
	Environment string
	DatabaseURL string
}

func NewConfig() *Config {
	return &Config{
		Environment: getEnvOrPanic("ENVIRONMENT"),
		DatabaseURL: getEnvOrPanic("DATABASE_URL"),
	}
}

func getEnvOrPanic(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic(fmt.Sprintf("required environment variable %q not set", key))
}
