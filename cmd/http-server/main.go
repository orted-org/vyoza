package main

import (
	"log"
	"net/http"
	"os"

	db "github.com/orted-org/vyoza/db/dao"
	configstore "github.com/orted-org/vyoza/internal/config_store"
	"github.com/orted-org/vyoza/internal/watcher"
)

type App struct {
	// db store
	store db.Store

	//logger
	logger *log.Logger

	// uptime and ssl watcher
	watcher *watcher.Watcher

	// config store
	configStore *configstore.Config

	// service quitter signal channel map
	quitters map[string]chan struct{}

	// channel for os signals
	osSignal chan os.Signal

	// http server
	srv *http.Server
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

	initServer(app)
	// go initWatcher(app)
	initConfigStore(app)
	go initCleaner(app)

	log.Fatal(app.srv.ListenAndServe())
}
