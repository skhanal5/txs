package repository

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/skhanal5/txs/internal/database/model"
	"go.uber.org/zap"
)

type AuthRepository interface {
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(user model.User) error
}

type PostgresAuthRepository struct {
	conn *pgxpool.Pool
	logger *zap.Logger
}

func NewPostgresAuthRepository(ctx context.Context, connection *pgxpool.Pool, logger *zap.Logger) *PostgresAuthRepository {
	return &PostgresAuthRepository{
		conn: connection,
		logger: logger,
	}
}

func (r *PostgresAuthRepository) GetUserByEmail(email string) (*model.User, error) {
	sql, params, err := goqu.From("users").Where(goqu.Ex{"email": email}).ToSQL()
	if err != nil {
		return nil, fmt.Errorf("failed to build select query: %w", err)
	}
	row := r.conn.QueryRow(context.Background(), sql, params...)
	user := &model.User{}
	err = row.Scan(&user.Email, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to execute select query: %w", err)
	}
	return user, nil
}

func (r *PostgresAuthRepository) CreateUser(user model.User) error {
	// TODO: unique email checks	
	
	sql, params, err := goqu.Insert("users").Rows(
		&user,
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
