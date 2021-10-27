package main

import (
	"context"
	"database/sql"
	"log"

	db "github.com/orted-org/vyoza/db/dao"
)

func initDB() (db.Store, error) {
	var err error
	tDB, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		return nil, err
	}
	q, err := db.Prepare(context.Background(), tDB)
	if err != nil {

		return nil, err
	}
	defer func() {
		err := tDB.Close()
		if err != nil {
			log.Fatal("could not close db connection", err)
		}
	}()
	return q, nil
}
