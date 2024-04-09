package main

import (
	"codejam.io/config"
	"codejam.io/database"
	"codejam.io/server"
)

func main() {
	// TODO setup logger

	config := new(config.Config)
	config.LoadFromFile("config.toml")

	database := new(database.Postgres)
	database.Initialize()

	server := server.Server{
		Config:   *config,
		Database: *database,
	}
	server.StartServer()
}
