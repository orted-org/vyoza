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
	authservice "github.com/orted-org/vyoza/internal/auth_service"
	configstore "github.com/orted-org/vyoza/internal/config_store"
	vault "github.com/orted-org/vyoza/internal/vault"
	watcher "github.com/orted-org/vyoza/internal/watcher"
	kvstore "github.com/orted-org/vyoza/pkg/kv_store"
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
			details, _ := app.store.GetUptimeWatchRequestByID(context.Background(), data.ID)
			conc, _ := app.store.GetUptimeResultStatsForID(context.Background(), data.ID)

			if conc.EndDate.Sub(conc.StartDate).Hours() >= float64(details.RetainDuration) {
				
				// creating the uptime conclusion
				var uc db.UptimeConclusion
				uc.UWRID = conc.UWRID
				uc.SuccessCount = conc.SuccessCount
				uc.WarningCount = conc.WarningCount
				uc.ErrorCount = conc.ErrorCount
				uc.MinResponseTime = conc.MinResponseTime
				uc.MaxResponseTime = conc.MaxResponseTime
				uc.AvgSuccessResponseTime = conc.AvgSuccessResponseTime
				uc.AvgWarningResponseTime = conc.AvgWarningResponseTime
				uc.StartDate = conc.StartDate
				uc.EndDate = conc.EndDate

				// deleting the older conclusion
				app.store.DeleteUptimeConclusionByUWRID(context.Background(), conc.UWRID)

				// storing the new conclusion
				app.store.AddUptimeConclusion(context.Background(), uc)

				// deleting the older uptime results
				app.store.DeleteUptimeResults(context.Background(), uc.UWRID)
			}
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

// vault init
func initVault(app *App) {
	app.vault = vault.New(app.store)
}

// config store init
func initConfigStore(app *App) {
	app.configStore = configstore.New(app.store)
}

//auth Service Init
func initAuthService(app *App) {
	app.authService = authservice.New(kvstore.New())
}
