package database

import (
	"codejam.io/config"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

type Postgres struct {
	Config config.Config
	Pool   *pgxpool.Pool
}

func (postgres *Postgres) Initialize() {
	pool, err := pgxpool.New(context.Background(), postgres.Config.Database.Url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unabled to create connection pool: %v\n", err)
		os.Exit(1)
	}
	pool.Config().MaxConns = postgres.Config.Database.MaxConnections
	postgres.Pool = pool
	fmt.Fprintf(os.Stdout, "Connected to database\n")
}
