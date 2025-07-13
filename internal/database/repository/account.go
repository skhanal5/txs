package repository

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/shopspring/decimal"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/skhanal5/txs/internal/database/model"
	"go.uber.org/zap"
)


type AccountRepository interface {
	GetAccountsById(userID string) ([]model.Account, error)
	CreateAccount(account model.Account) error
	TransferFunds(fromUser, toUser string, amount decimal.Decimal) error
}

type PostgresAccountRepository struct {
	conn   *pgxpool.Pool
	logger *zap.Logger
}


func NewPostgresAccountRepository(ctx context.Context, connection *pgxpool.Pool, logger *zap.Logger) *PostgresAccountRepository {
	return &PostgresAccountRepository{
		conn:   connection,
		logger: logger,
	}
}

func (r *PostgresAccountRepository) GetAccountsById(userID string) ([]model.Account, error) {
	sql, params, err := goqu.From("accounts").Select("*").Where(goqu.Ex{"user_id": "*"}).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to build select query: %w", err)
	}
	rows, err := r.conn.Query(context.Background(), sql, params...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute select query: %w", err)
	}
	defer rows.Close()
	accounts := []model.Account{}
	err = rows.Scan(&accounts)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal results: %w", err)
	}
	return accounts, nil
}

func (r *PostgresAccountRepository) CreateAccount(account model.Account) error {
	sql, params, err := goqu.Insert("accounts").Rows(
		&account,
	).ToSQL()

	if err != nil {
		return fmt.Errorf("failed to build insert query: %w", err)
	}

	_, err = r.conn.Exec(context.Background(), sql, params...)
	if err != nil {
		return fmt.Errorf("failed to execute insert query: %w", err)
	}
	return nil
}

func (r *PostgresAccountRepository) TransferFunds(fromUser, toUser string, amount decimal.Decimal) error {
	tx, err := r.conn.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(context.Background())


	sql, params, err := goqu.Update("accounts").Set(goqu.Record{"balance": goqu.L("balance - ?", amount)}).Where(goqu.Ex{"user_id": fromUser}).ToSQL()
	if err != nil {
		return fmt.Errorf("failed to build update query for sender: %w", err)
	}
	if _, err = tx.Exec(context.Background(), sql, params...); err != nil {
		return fmt.Errorf("failed to execute update query for sender: %w", err)
	}

	sql, params, err = goqu.Update("accounts").Set(goqu.Record{"balance": goqu.L("balance + ?", amount)}).Where(goqu.Ex{"user_id": toUser}).ToSQL()
	if err != nil {
		return fmt.Errorf("failed to build update query for receiver: %w", err)
	}
	if _, err = tx.Exec(context.Background(), sql, params...); err != nil {
		return fmt.Errorf("failed to execute update query for receiver: %w", err)
	}

	if err = tx.Commit(context.Background()); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}