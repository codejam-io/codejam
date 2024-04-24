package database

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type DBUser struct {
	Id            pgtype.UUID
	ServiceName   string
	ServiceUserId string
	DisplayName   string
	CreatedOn     pgtype.Timestamp
	ModifiedOn    pgtype.Timestamp
}

func CreateUser(serviceName string, serviceUserId string, serviceDisplayName string) *DBUser {
	conn, err := Pool.Acquire(context.Background())
	if err != nil {
		logger.Error("acquire conn error %v", err)
		return nil
	}
	defer conn.Release()

	row := conn.QueryRow(
		context.Background(),
		`INSERT INTO users (service_name, service_user_id, display_name)
		 VALUES ($1, $2, $3)
		 ON CONFLICT (service_name, service_user_id)
		 DO UPDATE
		 SET display_name = $3
		 RETURNING id`,
		serviceName, serviceUserId, serviceDisplayName)

	var result DBUser
	err = row.Scan(&result.Id)
	if err != nil {
		logger.Error("scan error %v", err)
	}

	return &result
}

func GetUser(userId pgtype.UUID) *DBUser {
	conn, err := Pool.Acquire(context.Background())
	if err != nil {
		logger.Error("acquire conn error %v", err)
		return nil
	}
	defer conn.Release()

	row := conn.QueryRow(
		context.Background(),
		`SELECT id, service_name, service_user_id, display_name, created_on
		 FROM users 
		 WHERE id = $1`,
		userId)

	var result DBUser
	err = row.Scan(&result.Id, &result.ServiceName, &result.ServiceUserId, &result.DisplayName, &result.CreatedOn)
	if err != nil && err != pgx.ErrNoRows {
		logger.Error("GetUser: scan error %v", err)
	}

	return &result
}
