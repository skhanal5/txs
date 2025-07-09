package handler

import (
	"net/http"
	"github.com/skhanal5/txs/internal/api/payload"
	"github.com/skhanal5/txs/internal/api/service"
	"go.uber.org/zap"
)

type AuthHandler struct {
	authService service.AuthService
	logger *zap.Logger
}

func NewAuthHandler(authService service.AuthService, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		logger: logger,
	}
}

func (a *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req payload.RegisterUserRequest
	if err := decode(&req, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := a.authService.RegisterUser(req); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (a *AuthHandler) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var req payload.AuthRequest
	if err := decode(&req, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	resp, err := a.authService.AuthenticateUser(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := encode(w, http.StatusOK, resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}