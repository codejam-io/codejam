package session

import (
	"context"
	"errors"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/gin-contrib/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
	"time"
)

type PGXStore struct {
	DBPool *pgxpool.Pool
	store  sessions.Store
}

type Session struct {
	ID         int64
	Key        string
	Data       string
	CreatedOn  time.Time
	ModifiedOn time.Time
	ExpiresOn  time.Time
}

func NewPGXStoreFromPool(pool *pgxpool.Pool) (*PGXStore, error) {
	pgxstore := &PGXStore{
		DBPool: pool,
	}
	return pgxstore, nil
}

func (pgxstore *PGXStore) Get(r *http.Request, name string) (*Session, error) {
	ctx := context.Background()
	conn, err := pgxstore.DBPool.Acquire(ctx)
	if err != nil {
		fmt.Printf("error aquiring pool connection: %+v\n", err)
		return nil, err
	}
	defer conn.Release()

	var sessions []*Session
	err = pgxscan.Select(ctx, conn, &sessions, "select * from http_sessions where key=$1", name)
	if err != nil {
		fmt.Printf("error reading session table: %+v\n", err)
		return nil, err
	}

	if len(sessions) == 0 {
		return nil, errors.New("no matching session record found")
	}
	return sessions[0], nil
}

// New should create and return a new session.
//
// Note that New should never return a nil session, even in the case of
// an error if using the Registry infrastructure to cache the session.
func (pgxstore *PGXStore) New(r *http.Request, name string) (*Session, error) {
	expiration, _ := time.ParseDuration("30d")
	return &Session{
		ID:         -1,
		Key:        "",
		Data:       "",
		CreatedOn:  time.Now(),
		ModifiedOn: time.Now(),
		ExpiresOn:  time.Now().Add(expiration),
	}, nil
}

// Save should persist session to the underlying store implementation.
func (pgxstore *PGXStore) Save(r *http.Request, w http.ResponseWriter, s *Session) error {
	ctx := context.Background()
	conn, err := pgxstore.DBPool.Acquire(ctx)
	if err != nil {
		fmt.Printf("error aquiring pool connection: %+v\n", err)
		return err
	}
	defer conn.Release()

	_, err = conn.Exec(ctx, "insert into http_sessions (key, data, created_on, modified_on, expires_on) values ($1, $2, $3, $4, $5)",
		s.Key,
		s.Data,
		s.CreatedOn,
		s.ModifiedOn,
		s.ExpiresOn)
	return err
}

func (pgxstore *PGXStore) Options(options sessions.Options) {

}
