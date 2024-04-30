package server

import (
	"codejam.io/database"
	"github.com/emicklei/pgtalk/convert"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (server *Server) GetAllEvents(ctx *gin.Context) {
	events, err := database.GetEvents()
	if err == nil {
		ctx.JSON(http.StatusOK, events)
	} else {
		ctx.Status(http.StatusInternalServerError)
	}
}

func (server *Server) GetEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	event, err := database.GetEvent(convert.StringToUUID(id))
	if err == nil {
		ctx.JSON(http.StatusOK, event)
	} else {
		logger.Error("GetEvent error: %v", err)
		ctx.Status(http.StatusInternalServerError)
	}
}

func (server *Server) CreateEvent(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userId := session.Get("userId")
	if userId != nil {
		event, err := database.CreateEvent(convert.StringToUUID(userId.(string)))
		if err == nil {
			ctx.JSON(http.StatusOK, event)
		} else {
			logger.Error("CreateEvent error: %v for user %s", err, userId)
			ctx.Status(http.StatusInternalServerError)
		}
	} else {
		ctx.Status(http.StatusUnauthorized)
	}
}

func (server *Server) UpdateEvent(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userId := session.Get("userId")
	if userId != nil {
		var event database.DBEvent
		err := ctx.ShouldBindJSON(&event)
		if err != nil {
			logger.Error("UpdateEvent Request ShouldBindJSON error: %v", err)
			ctx.Status(http.StatusBadRequest)
			return
		}
		event, err = database.UpdateEvent(event)
		if err != nil {
			logger.Error("Error calling database.UpdateEvent: %v", err)
			ctx.Status(http.StatusInternalServerError)
		} else {
			ctx.JSON(http.StatusOK, event)
		}
	} else {
		ctx.Status(http.StatusUnauthorized)
	}
}

func (server *Server) SetupEventRoutes() {
	group := server.Gin.Group("/event")
	{
		group.GET("/", server.GetAllEvents)
		group.GET("/:id", server.GetEvent)
		group.PUT("/:id", server.UpdateEvent)
		group.POST("/", server.CreateEvent)
	}
}
