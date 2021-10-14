package db

import (
	"context"
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
		log.Fatal(err)
		return
	}
	tq, err = Prepare(context.Background(), tDB)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer tq.Close()
	os.Exit(m.Run())
}
