package server

import (
	"codejam.io/database"
	"github.com/emicklei/pgtalk/convert"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) GetAllTeams(ctx *gin.Context) {
	teams, err := database.GetTeams()
	if err == nil {
		ctx.JSON(http.StatusOK, teams)
	} else {
		ctx.Status(http.StatusInternalServerError)
	}
}

func (server *Server) GetTeam(ctx *gin.Context) {
	id := ctx.Param("id")
	team, err := database.GetTeam(convert.StringToUUID(id))
	if err == nil {
		ctx.JSON(http.StatusOK, team)
	} else {
		logger.Error("GetEvent error: %v", err)
		ctx.Status(http.StatusInternalServerError)
	}
}

func (server *Server) CreateTeam(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userId := session.Get("userId")
	if userId != nil {
		team, err := database.CreateTeam(convert.StringToUUID(userId.(string)))
		if err == nil {
			ctx.JSON(http.StatusOK, team)
		} else {
			logger.Error("CreateTeam error: %v for user %s", err, userId)
			ctx.Status(http.StatusInternalServerError)
		}
	} else {
		ctx.Status(http.StatusUnauthorized)
	}
}

func (server *Server) UpdateTeam(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userId := session.Get("userId")
	if userId != nil {
		var team database.DBTeam
		err := ctx.ShouldBindJSON(&team)
		if err != nil {
			logger.Error("UpdateEvent Request ShouldBindJSON error: %v", err)
			ctx.Status(http.StatusBadRequest)
			return
		}
		team, err = database.UpdateTeam(team)
		if err != nil {
			logger.Error("Error calling database.UpdateEvent: %v", err)
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, team)
		}
	} else {
		ctx.Status(http.StatusUnauthorized)
	}
}

func (server *Server) SetupTeamRoutes() {
	group := server.Gin.Group("/team")
	{
		group.GET("/", server.GetAllTeams)
		group.GET("/:id", server.GetTeam)
		group.PUT("/:id", server.UpdateTeam)
		group.POST("/", server.CreateTeam)
	}
}
