package database

import (
	"codejam.io/config"
	"codejam.io/logging"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

var logger = logging.NewLogger(logging.Options{Name: "Database", Level: logging.INFO})

var Pool *pgxpool.Pool // concurrency safe

type Postgres struct {
	Config *config.DBConfig
	Pool   *pgxpool.Pool
}

func Initialize(config config.DBConfig) {
	logger.Info("Connecting to database")

	pool, err := pgxpool.New(context.Background(), config.Url)
	if err != nil {
		logger.Critical("Unable to initialise new pxpool.Pool: %+v\n", err)
		os.Exit(1)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		logger.Critical("Unable to connect to database: %+v\n", err)
		os.Exit(1)
	}

	Pool = pool
	logger.Info("Connected to database")
}
