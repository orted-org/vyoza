package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	db "github.com/orted-org/vyoza/db/dao"
	"github.com/orted-org/vyoza/internal/watcher"
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

		app.logger.Println("closing db connection")
		app.store.Close()

		// waiting to gracefully shutdown all the processes
		app.logger.Println("quitting application in 3s")
		time.Sleep(time.Second * 3)

		// finally shutting down the server
		app.logger.Println("shutting down the http server")
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		defer cancel()
		app.srv.Shutdown(ctx)
	}()
}

func initDB() (db.Store, error) {
	var err error

	// TODO: pass db info from config
	tDB, err := sql.Open("sqlite3", "../../db.db")
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

func initServer(app *App) {
	r := chi.NewRouter()
	initHandler(app, r)

	// TODO: send address from config file
	srv := http.Server{
		Addr:    "localhost:4000",
		Handler: r,
	}
	app.srv = &srv
}
func initWatcher(app *App) {

	// adding a quitter for uptime request watcher
	app.quitters["watcher"] = make(chan struct{})

	app.logger.Println("starting uptime request(ssl) watcher")

	i, err := app.store.GetAllUptimeWatchRequest(context.Background())
	if err != nil {
		app.logger.Println("could not get stored uptime watch requests", err.Error())
	} else {
		for _, v := range i {
			if v.Enabled {
				app.watcher.Register(watcher.WatcherParams{
					ID:              v.ID,
					Location:        v.Location,
					Interval:        v.Interval,
					ExpectedStatus:  v.ExpectedStatus,
					MaxResponseTime: v.MaxResponseTime,
				})
			}
			if v.SSLMonitor {
				app.watcher.RegisterSSL(watcher.SSLWatcherParams{
					ID:       v.ID,
					Location: v.Location,
					Interval: v.SSLInterval,
				})
			}
		}
	}

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
