package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type DBUser struct {
	Id            pgtype.UUID
	ServiceName   string
	ServiceUserId string
	DisplayName   string
	CreatedOn     pgtype.Timestamp
}

func CreateUser(serviceName string, serviceUserId string, serviceDisplayName string) DBUser {
	var user DBUser
	err := GetRow(
		&user,
		`INSERT INTO users (service_name, service_user_id, display_name)
		 VALUES ($1, $2, $3)
		 ON CONFLICT (service_name, service_user_id)
		 DO UPDATE
		 SET display_name = $3
		 RETURNING *`,
		serviceName, serviceUserId, serviceDisplayName)
	if err != nil {
		logger.Error("error getting user: %v", err)
	}
	return user
}

func GetUser(userId pgtype.UUID) DBUser {
	var user DBUser
	err := GetRow(
		&user,
		`SELECT *
		 FROM users 
		 WHERE id = $1`,
		userId)
	if err != nil {
		logger.Error("error getting user: %v", err)
	}
	return user
}
