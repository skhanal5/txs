package main

import (
	"context"
	"github.com/skhanal5/txs/internal/api/server"
	"github.com/skhanal5/txs/internal/config"
	"go.uber.org/zap"
)

func main() {
	cfg := config.NewConfig()
	logger := config.NewLogger("txs", cfg.Environment)
	ctx := context.Background()

	if err := server.Start(ctx, cfg, logger); err != nil {
		logger.Fatal("Server exited with error", zap.Error(err))
	}
}
