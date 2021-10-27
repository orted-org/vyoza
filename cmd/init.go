package main

import (
	"context"
	"database/sql"
	"os"
	"os/signal"
	"time"

	_ "github.com/mattn/go-sqlite3"
	db "github.com/orted-org/vyoza/db/dao"
)

// function to cleanup the open resources
func initCleaner(app *App) {
	app.osSignal = make(chan os.Signal, 1)
	signal.Notify(app.osSignal, os.Interrupt)
	go func() {
		<-app.osSignal
		app.logger.Println("starting cleaning up")

		app.logger.Println("removing all the go routines running services")
		for _, v := range app.quitters {
			v <- struct{}{}
		}

		// waiting to gracefully shutdown all the processes
		app.logger.Println("quitting application in 3s")
		time.Sleep(time.Second * 3)
		os.Exit(0)
	}()
}

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
