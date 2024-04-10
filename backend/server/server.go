package server

import (
	"codejam.io/config"
	"codejam.io/database"
	"codejam.io/logging"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var logger = logging.NewLogger(logging.Options{Name: "Server"})

type Server struct {
	Config   config.Config
	Database database.Postgres
	OAuth    *oauth2.Config
	Gin      *gin.Engine
}

func (server *Server) EmbedTest(context *gin.Context) {
	content, err := GetHtmlFile("html_files/index.html")
	if err != nil {
		context.Error(err)
		return
	}

	context.Writer.Write(content)
}

func (server *Server) SetupSessionStore() {
	// Todo - Redis? Postgres?
}

func (server *Server) StartServer() {
	logger.Info("Starting server...")

	server.Gin = gin.Default()

	server.OAuth = &oauth2.Config{
		ClientID:     server.Config.GitHub.Id,
		ClientSecret: server.Config.GitHub.Secret,
		Endpoint:     github.Endpoint,
		RedirectURL:  server.Config.GitHub.RedirectUrl,
	}

	server.SetupSessionStore()

	// Setup routes...
	server.SetupOAuthRoutes()

	// TODO: remove this test route
	server.Gin.GET("/", server.EmbedTest)

	// Start the server...
	server.Gin.Run(server.Config.Server.Listen)
}
