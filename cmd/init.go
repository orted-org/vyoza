package main

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
	db "github.com/orted-org/vyoza/db/dao"
)

func initDB() (db.Store, error) {
	var err error
	
	// TODO: pass db info from config
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

func initWatcher(app *App) {

	// adding a quitter for uptime request watcher
	app.quitters["watcher"] = make(chan struct{})

	app.logger.Println("starting uptime request(ssl) watcher")
	// TODO: add logic to send data to web socket clients
	for {
		select {
		case data := <-app.watcher.Result:
			app.store.AddUptimeResult(context.Background(), db.AddUptimeResultParams{
				UWRID:        data.ID,
				ResponseTime: data.ResponseTime,
				Remark:       data.Remark,
			})
		case sslData := <-app.watcher.SSLResult:
			app.store.UpdateUptimeSSLInfoByUWRID(context.Background(), db.UptimeSSLInfo{
				UWRID:      sslData.ID,
				IsValid:    sslData.IsValid,
				ExpiryDate: sslData.ExpiryDate,
				Remark:     sslData.Remark,
				UpdatedAt:  time.Now().UTC(),
			})
		case <-app.quitters["watcher"]:
			app.logger.Println("quitting uptime request(ssl) watcher")
		}
	}
}
