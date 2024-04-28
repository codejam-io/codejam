package database

import "github.com/jackc/pgx/v5/pgtype"

// These fields need to match the order of the fields in the database
type DBEvent struct {
	Id              pgtype.UUID
	StatusId        int
	Title           string
	Description     string
	Rules           string
	OrganizerUserId pgtype.UUID
	MaxTeams        int
	StartsAt        pgtype.Timestamp
	EndsAt          pgtype.Timestamp
	CreatedOn       pgtype.Timestamp
}

func CreateEvent(organizerUserId pgtype.UUID) (DBEvent, error) {
	var event DBEvent
	err := GetRow(
		&event,
		`INSERT INTO events
           (status_id, title, description, rules, organizer_user_id)
         VALUES
           ((SELECT id FROM statuses WHERE code = 'PLANNING'),
            '',
            '',
            '',
            $1)
         RETURNING *
         `,
		organizerUserId)
	return event, err
}

func GetEvent(eventId pgtype.UUID) (DBEvent, error) {
	var event DBEvent
	err := GetRow(
		&event,
		`SELECT * FROM events WHERE id = $1`,
		eventId)
	return event, err
}
