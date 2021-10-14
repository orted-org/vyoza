package db

import (
	"context"
	"database/sql"
)

type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func NewDB(db DB) *Queries {
	return &Queries{db: db}
}

type Queries struct {
	db DB
}
