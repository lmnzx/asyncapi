package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lmnzx/asyncapi/api"
	"github.com/lmnzx/asyncapi/config"
	"github.com/lmnzx/asyncapi/store"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg, err := config.New()
	if err != nil {
		return err
	}

	dbpool, err := pgxpool.New(ctx, cfg.DatabaseConnectionString)
	if err != nil {
		return err
	}

	if err := dbpool.Ping(ctx); err != nil {
		return err
	}

	store := store.New(dbpool)

	api.New(cfg, store).Start(ctx)

	return nil
}
