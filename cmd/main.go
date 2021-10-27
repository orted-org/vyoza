package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	db "github.com/orted-org/vyoza/db/dao"
	"github.com/orted-org/vyoza/internal/watcher"
)

type App struct {
	// db store
	store db.Store

	//logger
	logger *log.Logger

	// uptime and ssl watcher
	watcher *watcher.Watcher

	// service quitter signal channel map
	quitters map[string]chan struct{}

	osSignal chan os.Signal
}

var (
	lo = log.New(os.Stdout, "",
		log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {
	store, err := initDB()
	if err != nil {
		log.Fatal("error initializing db store", err)
		return
	}
	app := &App{
		store:    store,
		watcher:  watcher.New(),
		quitters: make(map[string]chan struct{}),
		logger:   lo,
	}

	r := chi.NewRouter()
	initHandler(app, r)
	initWatcher(app)

	// TODO: server handling
	srv := http.Server{
		Addr:    "localhost:4000",
		Handler: r,
	}

	log.Fatal(srv.ListenAndServe())
}
