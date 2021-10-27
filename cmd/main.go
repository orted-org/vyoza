package main

import (
	db "github.com/orted-org/vyoza/db/dao"
	"github.com/orted-org/vyoza/internal/watcher"
)

type App struct {
	// db store
	store *db.Store

	// uptime and ssl watcher
	watcher watcher.Watcher
}

func main() {

}
