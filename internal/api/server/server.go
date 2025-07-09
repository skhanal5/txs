package server

import (
	"context"
	"net/http"
	"time"

	"github.com/skhanal5/txs/internal/api/handler"
	"github.com/skhanal5/txs/internal/api/service"
	"github.com/skhanal5/txs/internal/config"
	"github.com/skhanal5/txs/internal/database/postgres"
	"github.com/skhanal5/txs/internal/database/repository"
	"go.uber.org/zap"
)

func Start(ctx context.Context, cfg *config.Config, logger *zap.Logger) error {

	db, err := postgres.NewConnection(ctx, cfg.DatabaseURL)
	if err != nil {
		logger.Fatal("Failed to connect to database", zap.Error(err))
	}
	defer db.Close()

	authRepo := repository.NewPostgresAuthRepository(ctx, db, logger)
	authService := service.NewAuthService(authRepo, logger)
	authHandler := handler.NewAuthHandler(authService, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /auth/register", authHandler.RegisterUser)
	mux.HandleFunc("POST /auth/login", authHandler.AuthenticateUser)
	logger.Info("Server starting on :8080")

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,	
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
	return server.ListenAndServe()
}