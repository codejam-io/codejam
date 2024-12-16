package database

import (
	"embed"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/stdlib"
	"os"
)

//go:embed migrations/*
var migrationsFS embed.FS

func RunMigrations() {
	logger.Info("Running DB Migrations...")
	stdConn := stdlib.OpenDBFromPool(Pool)

	migrationConfig := pgx.Config{}
	sqlDriver, err := pgx.WithInstance(stdConn, &migrationConfig)
	if err != nil {
		logger.Critical("Error setting up migrations DB connection: %v", err)
		os.Exit(1)
	}

	sourceDriver, err := iofs.New(migrationsFS, "migrations")
	if err != nil {
		logger.Critical("Error setting up migrations FS: %v", err)
		os.Exit(1)
	}

	m, err := migrate.NewWithInstance("embeddedFS", sourceDriver, "postgres", sqlDriver)
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		logger.Critical("Migrations failure: %v", err)
		os.Exit(1)
	}
	logger.Info("Migrations completed")
}
