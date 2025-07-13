package config

import (
	"fmt"
	"os"
)

type Config struct {
	Environment string
	DatabaseURL string
	JWTSecret []byte
}

func NewConfig() *Config {
	return &Config{
		Environment: getEnvAsStringOrPanic("ENVIRONMENT"),
		DatabaseURL: getEnvAsStringOrPanic("DATABASE_URL"),
		JWTSecret: getEnvAsByteOrPanic("JWT_SECRET"),
	}
}

func getEnvAsByteOrPanic(key string) []byte {
	if value, exists := os.LookupEnv(key); exists {
		return []byte(value)
	}
	panic(fmt.Sprintf("required environment variable %q not set", key))
}

func getEnvAsStringOrPanic(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	panic(fmt.Sprintf("required environment variable %q not set", key))
}
