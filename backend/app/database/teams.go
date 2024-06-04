package database

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type DBTeam struct {
	Id              pgtype.UUID      `db:"id"`
	EventId         pgtype.UUID      `db:"event_id"`
	OwnerUserId     pgtype.UUID      `db:"owner_user_id" json:"ownerUserId-hidden"`
	Name            string           `db:"name"`
	Visibility      string           `db:"visibility"`
	Timezone        string           `db:"timezone"`
	Technologies    string           `db:"technologies"`
	Availability    string           `db:"availability"`
	Description 	string 			 `db:"description"`
	CreatedOn       pgtype.Timestamp `db:"created_on" json:"createdOn-hidden"`
	OwnerName	    *OwnerName		 `db:"display_name, omitempty" json:"-"` // this references the User struct
}

type OwnerName struct {
    DisplayName string `db:"display_name, omitempty"`
}

func CreateTeam(team DBTeam) (DBTeam, error) {
	team, err := GetRow[DBTeam](
		`INSERT INTO teams
            (event_id, owner_user_id, name, visibility, timezone, technologies, availability, description)
            VALUES
			($1, $2, $3, $4, $5, $6, $7, $8)
        RETURNING id, event_id, owner_user_id, name, visibility, timezone, technologies, availability, description, created_on
		`,
		team.EventId, team.OwnerUserId, team.Name, team.Visibility, team.Timezone, team.Technologies, team.Availability, team.Description)
	return team, err
}

// stepp 5: api get team info
func GetTeam(teamId pgtype.UUID) (DBTeam, error) {
	team, err := GetRow[DBTeam](
		`SELECT 
			teams.id,
			teams.event_id,
			teams.owner_user_id,
			teams.name,
			teams.visibility,
			teams.timezone,
			teams.technologies,
			teams.availability,
			teams.description,
			teams.created_on,
			users.display_name
		FROM teams
		LEFT JOIN users ON teams.owner_user_id = users.id
		WHERE teams.id = $1`,
		teamId)
		// `SELECT * FROM teams WHERE id = $1`,
		// teamId)
	if err!= nil {
		return DBTeam{}, err
	}
	return team, nil
}

func GetTeams() ([]DBTeam, error) {
	result, err := GetRows[DBTeam](`SELECT * FROM teams`)
	return result, err
}

func UpdateTeam(team DBTeam) (DBTeam, error) {
	event, err := GetRow[DBTeam](
		`UPDATE teams
            SET name=$2,
                visibility=$3,
				timezone=$4,
				technologies=$5,
				availability=$6,
				description=$7,
		WHERE id=$1
		RETURNING *`,
		team.Id, team.Name, team.Visibility, team.Timezone, team.Technologies, team.Availability, team.Description)
	return event, err
}
