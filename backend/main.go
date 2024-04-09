package main

import (
	"codejam.io/config"
	"codejam.io/database"
	"codejam.io/server"
	"github.com/gin-gonic/gin"
)

func main() {
	// TODO setup logger

	// Disables the debug logging... Comment this out to enable debug logging for GIN
	gin.SetMode(gin.ReleaseMode)

	config := new(config.Config)
	config.LoadFromFile("config.toml")

	database := database.Postgres{Config: &config.Database}
	database.Initialize()

	server := server.Server{
		Config:   *config,
		Database: database,
	}
	server.StartServer()
}
