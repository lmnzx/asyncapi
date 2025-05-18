package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

var (
	ctx                      = context.Background()
	PostgresConnectionString string
	DBPool                   *pgxpool.Pool
)

func setupContainer() error {
	dbname := fmt.Sprintf("testdb_%d", time.Now().UnixNano())

	postgresContainer, err := postgres.Run(ctx,
		"postgres:17-alpine",
		postgres.WithInitScripts("../schema.sql"),
		postgres.WithDatabase(dbname),
		postgres.WithUsername("test_user"),
		postgres.WithPassword("test_password"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(30*time.Second),
		),
	)
	if err != nil {
		return err
	}

	postgresHost, err := postgresContainer.Host(ctx)
	if err != nil {
		return err
	}

	postgresPort, err := postgresContainer.MappedPort(ctx, "5432/tcp")
	if err != nil {
		return err
	}

	PostgresConnectionString = fmt.Sprintf("postgres://test_user:test_password@%s:%s/%s?sslmode=disable", postgresHost, postgresPort.Port(), dbname)

	return nil
}

func TestMain(m *testing.M) {
	if err := setupContainer(); err != nil {
		log.Fatal(err.Error())
	}

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestConnection(t *testing.T) {
	dbpool, err := pgxpool.New(ctx, PostgresConnectionString)
	if err != nil {
		t.Error(err)
	}

	if err := dbpool.Ping(ctx); err != nil {
		t.Error(err)
	}
	DBPool = dbpool
	t.Log("Database connection is OK")
}
