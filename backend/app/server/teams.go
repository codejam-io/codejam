package server

import (
	"fmt"
	"net/http"

	"codejam.io/database"
	"github.com/emicklei/pgtalk/convert"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type CreateTeamRequest struct {
	// sent from UI, to be processed in the server into a DBTeam structure
	// referenced in server/teams.go/CreateTeam line 53
	EventId 		string 
	Name			string			
	Visibility		string
	Availability 	string
	Description		string
	Technologies 	string
	Timezone 		string
}


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
	// ctx argument with type *gin.Context has the HTTP post information. 
	// Step 4: Post Team Data API
	session := sessions.Default(ctx)
	userId := session.Get("userId")
	if userId != nil {
		// declare type team

		var team database.DBTeam
		var teamReq CreateTeamRequest
		//event_id = convert.StringToUUID(ctx.EventId.string())	
		
		// shouldbindJSON binds the POST-req-JSON-info to the provided structure in ()
		// err should be <nil> if there are no errors (this is a context feature), otherwise -> line 60
		err := ctx.ShouldBindJSON(&teamReq)
		fmt.Println(teamReq, "==", err)

		if err != nil {
			logger.Error("CreateTeam Request ShouldBindJSON error: %v", err)
			ctx.Status(http.StatusBadRequest)
			return
		}

		// CONVERT teamReq to team
		team.OwnerUserId = convert.StringToUUID(userId.(string))
		team.EventId = convert.StringToUUID(teamReq.EventId)
		team.Name = teamReq.Name
		team.Availability = teamReq.Availability
		team.Description = teamReq.Description
		team.Visibility = teamReq.Visibility
		team.Technologies = teamReq.Technologies
		team.Timezone = teamReq.Timezone

		fmt.Printf("%+v",teamReq) 
		//{EventId:28247e11-319e-4f56-8200-3e6366b79f04 Name:jarvie Visibility: Availability:sometimes Description:im shy}[GIN]

		team, err = database.CreateTeam(team)
		fmt.Printf("%+v", team)

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
		// Step 3: Post Team Data API
		group.POST("/", server.CreateTeam)
	}
}
