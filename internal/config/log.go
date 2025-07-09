package config

import (
	"go.uber.org/zap"
)


func NewLogger(service string, environment string) (*zap.Logger) {
	logger, _ := zap.NewProduction(zap.Fields(
		zap.String("env", environment),
		zap.String("service", service),
	))

	if environment == "" || environment == "development" {
		logger, _ = zap.NewDevelopment()
	}
	return logger
}