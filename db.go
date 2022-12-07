package saaswhip

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Datastorer is an interface for working with the Database
type Datastorer interface {
	// Ping pings the DB pool.
	Ping(ctx context.Context) error
	// BeginTx starts a pgx.Tx using the input context
	BeginTx(ctx context.Context) (pgx.Tx, error)
	// RollbackTx rolls back the input pgx.Tx
	RollbackTx(ctx context.Context, tx pgx.Tx, err error) error
	// CommitTx commits the Tx
	CommitTx(ctx context.Context, tx pgx.Tx) error
}

// DBTX interface mirrors the interface generated by https://github.com/kyleconroy/sqlc
// to allow passing a Pool or a Tx
type DBTX interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
}

// PingServicer pings the database and responds whether it is up or down
type PingServicer interface {
	Ping(ctx context.Context, lgr zerolog.Logger) PingResponse
}

// PingResponse is the response struct for the PingService
type PingResponse struct {
	DBUp bool `json:"db_up"`
}

// NewNullString returns a null if s is empty, otherwise it returns
// the string which was input
func NewNullString(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

// NewNullTime returns a null if t is the zero value for time.Time,
// otherwise it returns the time which was input
func NewNullTime(t time.Time) sql.NullTime {
	if t.IsZero() {
		return sql.NullTime{}
	}
	return sql.NullTime{
		Time:  t,
		Valid: true,
	}
}

// NewNullInt64 returns a null if i == 0, otherwise it returns
// the int64 which was input.
func NewNullInt64(i int64) sql.NullInt64 {
	if i == 0 {
		return sql.NullInt64{}
	}
	return sql.NullInt64{
		Int64: i,
		Valid: true,
	}
}

// NewNullInt32 returns a null if i == 0, otherwise it returns
// the int32 which was input.
func NewNullInt32(i int32) sql.NullInt32 {
	if i == 0 {
		return sql.NullInt32{}
	}
	return sql.NullInt32{
		Int32: i,
		Valid: true,
	}
}

// NewNullUUID returns a null if i == uuid.Nil, otherwise it returns
// the int32 which was input.
func NewNullUUID(i uuid.UUID) uuid.NullUUID {
	if i == uuid.Nil {
		return uuid.NullUUID{}
	}
	return uuid.NullUUID{
		UUID:  i,
		Valid: true,
	}
}
