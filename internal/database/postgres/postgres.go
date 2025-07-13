package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxdecimal "github.com/jackc/pgx-shopspring-decimal"
)

func NewConnection(ctx context.Context, url string) (*pgxpool.Pool, error) {

	cfg, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("unable to parse database URL: %v", err)
	}

	
	cfg.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		pgxdecimal.Register(conn.TypeMap())
		return nil
	}

	conn, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}
	return conn, nil
}
