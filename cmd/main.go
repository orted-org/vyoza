package main

import (
	"log"

	"github.com/go-chi/chi/v5"
	db "github.com/orted-org/vyoza/db/dao"
	"github.com/orted-org/vyoza/internal/watcher"
)

type App struct {
	// db store
	store db.Store

	// uptime and ssl watcher
	watcher watcher.Watcher
}

func main() {
	store, err := initDB()
	if err != nil {
		log.Fatal("error initializing db store", err)
		return
	}
	app := &App{
		store: store,
	}

	r := chi.NewRouter()
	initHandler(app, r)

}
