package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

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

func (server *Server) SetupOAuthRoutes() {
	logger.Info("Setting up OAuth routes...")
	
	group := server.Gin.Group("/oauth")
	{
		group.GET("/redirect", server.GetOAuthRedirect)
		group.GET("/callback", server.GetOAuthCallback)
	}
}
