package server

import (
	"codejam.io/database"
	"github.com/emicklei/pgtalk/convert"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) GetUser(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userId := session.Get("userId")
	if userId != nil {
		dbUser := database.GetUser(convert.StringToUUID(userId.(string)))
		ctx.JSON(http.StatusOK, dbUser)
		return
	}
	ctx.Status(http.StatusUnauthorized)
}

// Logout is a GET route for logging out a user.
// This involved clearing the session cookie, clearing/deleting the entry from the session store
func (server *Server) Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Clear()
	err := session.Save()
	if err != nil {
		logger.Error("Logout: error saving session: %v", err)
	}

	// Have to manually do this, not sure why the session middleware doesn't handle this...
	ctx.SetCookie(SessionCookieName, "", -1, "/", "", false, false)
	ctx.Redirect(http.StatusFound, ctx.Request.Header.Get("Referer"))
}

func (server *Server) SetupUserRoutes() {
	logger.Info("Setting up User routes...")

	group := server.Gin.Group("/user")
	{
		group.GET("/", server.GetUser)
		group.GET("/logout", server.Logout)
	}
}
