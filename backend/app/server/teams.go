package server

import (
	"fmt"
	"net/http"

	"codejam.io/database"
	"github.com/emicklei/pgtalk/convert"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateTeamRequest struct {
	// sent from UI, to be processed in the server into a DBTeam structure
	// referenced in server/teams.go/CreateTeam line 53
	EventId      string
	Name         string
	Visibility   string
	Availability string
	Description  string
	Technologies string
	Timezone     string
}

type GetTeamResponse struct {
	Team    *database.DBTeam
	Event   *database.DBEvent
	Members *[]database.DBTeamMemberInfo // array(slice) of a struct
}

func (server *Server) GetAllTeams(ctx *gin.Context) {
	teams, err := database.GetTeams()
	if err == nil {
		ctx.JSON(http.StatusOK, teams)
	} else {
		ctx.Status(http.StatusInternalServerError)
	}
}

// stepp 4: GET team info
// purpose is to construct the DBTeamMemberInfo
func (server *Server) GetTeamInfo(ctx *gin.Context) {
	id := convert.StringToUUID(ctx.Param("id"))

	var teamResponse GetTeamResponse
	var team database.DBTeam
	var event database.DBEvent
	var members *[]database.DBTeamMemberInfo //user info based on teamId

	team, err := database.GetTeam(id)
	if err != nil {
		logger.Error("failed to get team: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get team: %v", err)})
		return
	}

	event, err = database.GetEvent(team.EventId)
	if err != nil {
		logger.Error("failed to get event: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get event: %v", err)})
		return
	}

	members, err = database.GetMembersByTeamId(team.Id)
	if err != nil {
		logger.Error("failed to get event: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to get members: %v", err)})
		return
	}

	// attach all 3 structures to GetTeamResponse --> nested structs turn into nested JSON (with ctx.JSON)
	teamResponse.Team = &team
	teamResponse.Event = &event
	teamResponse.Members = members

	ctx.JSON(http.StatusOK, teamResponse)
}

func (server *Server) CreateTeam(ctx *gin.Context) {
	// ctx of *gin.Context has HTTP request info.
	// Step 4: Post Team Data API (TWO PARTS 1) create team 2) add team members)
	session := sessions.Default(ctx)
	userId := session.Get("userId")
	if userId != nil {
		ctx.Status(http.StatusUnauthorized)
	}
	strUserId := userId.(string)

	var team database.DBTeam
	var teamReq CreateTeamRequest
	// var tempMember CreateTeamMember

	// shouldbindJSON binds the POST-req-JSON-info to the provided structure in ()
	// err should be <nil> (ctx feature)
	err := ctx.ShouldBindJSON(&teamReq)
	if err != nil {
		logger.Error("CreateTeam Request ShouldBindJSON error: %v", err)
		ctx.Status(http.StatusBadRequest)
		return
	}

	// CONVERT teamReq to team
	team.EventId = convert.StringToUUID(teamReq.EventId)
	team.Name = teamReq.Name
	team.Availability = teamReq.Availability
	team.Description = teamReq.Description
	team.Visibility = teamReq.Visibility
	team.Technologies = teamReq.Technologies
	team.Timezone = teamReq.Timezone

	fmt.Printf("%+v", team)
	// INSERTS TEAM into DB
	// PART 1/2 DONE
	teamUUID, err := database.CreateTeam(team)
	if err != nil {
		logger.Error("Error trying to CreateTeam(team)")
		ctx.Status(http.StatusBadRequest)
		return
	}
	// PART 2/2 DONE
	// construct TeamMember
	// tempMember.Role = "owner"
	// tempMember.TeamId = teamUUID
	// tempMember.UserID = convert.StringToUUID(strUserId)
	_, err = database.AddTeamMember(convert.StringToUUID(strUserId), teamUUID, "owner")

	if err == nil {
		fmt.Println("Successfully added team member")
		ctx.JSON(http.StatusCreated, map[string]pgtype.UUID{
			"id": teamUUID,
		})
	} else {
		fmt.Println(err)
		logger.Error("AddTeamMember error: %v for user %s", err, strUserId)
		ctx.Status(http.StatusInternalServerError)
		return
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
		group.POST("/", server.CreateTeam)
		group.GET("/", server.GetAllTeams)
		group.GET("/:id", server.GetTeamInfo)
		// group.PUT("/:id", server.UpdateTeam)
		// Step 3: Post Team Data API
	}
}
