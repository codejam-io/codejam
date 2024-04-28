package database

import (
	"codejam.io/config"
	"codejam.io/logging"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"reflect"
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

func GetRow[T any](output *T, query string, args ...any) error {
	conn, err := Pool.Acquire(context.Background())
	if err != nil {
		logger.Error("acquire conn error %v", err)
		return err
	}
	defer conn.Release()

	// Look at the output value passed in, and get pointers to all the individual fields.  Those will be the Scan params
	var pointers []any
	v := reflect.ValueOf(output).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.CanAddr() {
			pointers = append(pointers, field.Addr().Interface())
		}
	}

	row := conn.QueryRow(context.Background(),
		query,
		args...)
	err = row.Scan(pointers...)
	if err != nil {
		logger.Error("Error scanning row results: %v", err)
		return err
	}

	return nil
}
