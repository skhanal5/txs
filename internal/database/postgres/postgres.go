package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

func buildConnectionString(username string, password string, port string, database string) string {
	connectionString := "postgres://%s:%s@localhost:%s/%s"
	return fmt.Sprintf(connectionString, username, password, port, database)
}	

func NewConnection(ctx context.Context, username string, password string, port string, database string) (*pgxpool.Pool, error) {
	url := buildConnectionString(username, password, port, database)
	conn, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)		
	}
	return conn, nil
}	