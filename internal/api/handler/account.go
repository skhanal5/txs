package handler

import (
	"net/http"

	"github.com/skhanal5/txs/internal/api/payload"
	"github.com/skhanal5/txs/internal/api/service"
	"go.uber.org/zap"
)

type AccountHandler struct {
	accountService service.AccountService
	logger      *zap.Logger
}

func NewAccountHandler(accountService service.AccountService, logger *zap.Logger) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
		logger:      logger,
	}
}

func (h *AccountHandler) GetAccountsById(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	accounts, err := h.accountService.GetAccountsById(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp := payload.AccountsResponse{Accounts: accounts}
	if err := encode(w, http.StatusOK, resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}