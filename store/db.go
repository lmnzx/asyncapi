package store

import (
	"context"
	"fmt"
	"time"

	"github.com/lmnzx/asyncapi/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPgDb(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, cfg.DatabaseConnectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}
	if err := dbpool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database connection: %w", err)
	}

	return dbpool, nil
}
