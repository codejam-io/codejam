package server

import (
	"codejam.io/config"
	"codejam.io/database"
	"codejam.io/logging"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"os"
)

var logger = logging.NewLogger(logging.Options{Name: "Server", Level: logging.DEBUG})

type Server struct {
	Config   config.Config
	Database database.Postgres
	OAuth    *oauth2.Config
	Gin      *gin.Engine
}

func (server *Server) SetupSessionStore() {
	store, err := redis.NewStore(
		server.Config.Redis.Size,
		server.Config.Redis.Protocol,
		server.Config.Redis.Address,
		server.Config.Redis.Password,
		[]byte(""))

	if err != nil {
		logger.Critical("error initializing Redis session store: %v", err)
		os.Exit(1)
	}
	server.Gin.Use(sessions.Sessions("session", store))
}

func (server *Server) StartServer() {
	logger.Info("Starting server...")

	server.Gin = gin.Default()

	server.SetupOAuth()
	server.SetupSessionStore()

	// Setup routes...
	server.SetupOAuthRoutes()
	server.SetupStaticRoutes()

	// Start the server...
	server.Gin.Run(server.Config.Server.Listen)
}
