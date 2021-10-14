package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

const (
	dbSource = "../../db.db"
)

var tq *Queries

func TestMain(m *testing.M) {
	var err error
	tDB, err := sql.Open("sqlite3", dbSource)
	if err != nil {
		log.Fatal("could not connect to db")
	}
	tq = NewDB(tDB)

	os.Exit(m.Run())
}
