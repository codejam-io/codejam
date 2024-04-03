package server

import (
	"codejam.io/config"
	"codejam.io/database"
	"codejam.io/session"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Config       config.Config
	Database     database.Postgres
	OAuth        *oauth2.Config
	Gin          *gin.Engine
	SessionStore session.PGXStore
}

func (server *Server) EmbedTest(context *gin.Context) {
	content, err := GetHtmlFile("html_files/index.html")
	if err != nil {
		context.Error(err)
		return
	}

	context.Writer.Write(content)
}

func (server *Server) GetOAuthRedirect(context *gin.Context) {
	url := server.OAuth.AuthCodeURL("state")
	context.Redirect(http.StatusFound, url)
}

func (server *Server) GetOAuthCallback(context *gin.Context) {
	authCode := context.Query("code")
	token, err := server.OAuth.Exchange(oauth2.NoContext, authCode)
	if err != nil {
		log.Fatal(err)
	}
	client := server.OAuth.Client(oauth2.NoContext, token)
	fmt.Printf("DEBUG: OAuth Callback result: %+v\n", client)

	// TODO redirect somewhere
}

func (server *Server) SetupSessionStore() {
	_, err := session.NewPGXStoreFromPool(server.Database.Pool)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to setup session store: %v\n", err)
		os.Exit(1)
	}
	// TODO setup session store in Gin
}

func (server *Server) StartServer() {
	server.Gin = gin.Default()

	server.OAuth = &oauth2.Config{
		ClientID:     server.Config.GitHub.Id,
		ClientSecret: server.Config.GitHub.Secret,
		Endpoint:     github.Endpoint,
		RedirectURL:  server.Config.GitHub.RedirectUrl,
	}

	server.SetupSessionStore()

	server.Gin.GET("/", server.EmbedTest)
	server.Gin.GET("/oauth/redirect/", server.GetOAuthRedirect)
	server.Gin.GET("/oauth/callback/", server.GetOAuthCallback)
	server.Gin.Run(server.Config.Server.Listen)
}
