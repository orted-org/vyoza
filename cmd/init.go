package main

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	db "github.com/orted-org/vyoza/db/dao"
)

func initDB() (db.Store, error) {
	var err error
	tDB, err := sql.Open("sqlite3", "../db.db")
	if err != nil {
		return nil, err
	}
	q, err := db.Prepare(context.Background(), tDB)
	if err != nil {
		return nil, err
	}

	// TODO: closing db connection
	return q, nil
}
