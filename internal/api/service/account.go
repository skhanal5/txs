package service

import (
	"github.com/shopspring/decimal"
	"github.com/skhanal5/txs/internal/api/payload"
	"github.com/skhanal5/txs/internal/database/model"
	"github.com/skhanal5/txs/internal/database/repository"
	"go.uber.org/zap"
)

type AccountService interface {
	GetAccountsById(userID string) ([]payload.Account, error)
	CreateAccount(account payload.CreateAccountRequest) error
	TransferFunds(fromUser, toUser string, amount decimal.Decimal) error
}

type accountService struct {
	repository repository.AccountRepository
	logger     *zap.Logger
}

func NewAccountService(accountRepository repository.AccountRepository, logger *zap.Logger) AccountService {
	return &accountService{
		repository: accountRepository,
		logger:     logger,
	}
}

func fromPayloadToModel(account payload.Account) (model.Account, error) {
	convertedBalance, err := decimal.NewFromString(account.Balance)
	if err != nil {
		return model.Account{}, err
	}

	return model.Account{
		UserId:        account.UserID,
		Balance:       convertedBalance,
		CurrencyCode:  account.CurrencyCode,
		Status:        account.Status,
		Type:          account.Type,
		AccountNumber: account.AccountNumber,
	}, nil
}

func fromModelToPayload(account model.Account) payload.Account {
	return payload.Account{
		UserID:        account.UserId,
		Balance:       account.Balance.String(),
		CurrencyCode:  account.CurrencyCode,
		Status:        account.Status,
		Type:          account.Type,
		AccountNumber: account.AccountNumber,
	}
}


func (s *accountService) GetAccountsById(userID string) ([]payload.Account, error) {
	accounts, err := s.repository.GetAccountsById(userID)
	if err != nil {
		s.logger.Error("failed to get accounts by user ID", zap.Error(err))
		return nil, err
	}
	var accountPayloads []payload.Account
	for _, account := range accounts {
		payloadAccount := fromModelToPayload(account)
		accountPayloads = append(accountPayloads, payloadAccount)
	}
	return accountPayloads, nil
}

func (s *accountService) CreateAccount(account payload.CreateAccountRequest) error {
	payload := payload.Account(account)
	modelAccount, err := fromPayloadToModel(payload)
	if err != nil {
		s.logger.Error("failed to convert payload to model", zap.Error(err))
		return err
	}
	err = s.repository.CreateAccount(modelAccount)
	if err != nil {
		s.logger.Error("failed to create account", zap.Error(err))
		return err
	}
	return nil
}

func (s *accountService) TransferFunds(fromUser, toUser string, amount decimal.Decimal) error {
	err := s.repository.TransferFunds(fromUser, toUser, amount)
	if err != nil {
		s.logger.Error("failed to transfer funds", zap.Error(err))
		return err
	}
	return nil
}