package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("insufficient arguments")
		return
	}
	mode := os.Args[1]
	if !(mode == "up" || mode == "down") {
		log.Fatal("invalid migration mode")
		return
	}
	version := os.Args[2]

	fileByte, err := ioutil.ReadFile(filepath.Join(".", "db", "migration", fmt.Sprintf("v%s.%s.sql", version, mode)))
	if err != nil {
		log.Fatal(err)
	}

	if len(fileByte) == 0 {
		log.Fatal("empty file")
	}

	db, err := sql.Open("sqlite3", "db.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(fileByte))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Migrated", mode, version)
}
