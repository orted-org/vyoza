package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func middlewareOne(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		x := true
		if x {
			w.Write([]byte("auth failed"))
			return
		}
		next.ServeHTTP(w, r)
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	w.Write([]byte("OK"))
}

func initHandler(app *App, r *chi.Mux) {
  
  // uptime
	r.Post("/uptime", app.handleCreateWatchReq)
	r.Get("/uptime", app.handleGetWatchReq)
	r.Put("/uptime/{id}", app.handleUpdateWatchReq)
	r.Delete("/uptime/{id}", app.handleDeleteWatchReq)

	// vault
	r.Get("/vault/{name}", app.handleGetVault)
	r.Post("/vault", app.handleSetVault)
	r.Put("/vault", app.handleUpdateVault)
	r.Delete("/vault/{name}", app.handleDeleteVault)

	// config store
	r.Get("/cs/{name}", app.handleGetConfig)
	r.Post("/cs", app.handleSetConfig)
	r.Delete("/cs/{name}", app.handleDeleteConfig)


	//just chscking
	finalHandler := http.HandlerFunc(final)
	r.Get("/protected",middlewareOne(finalHandler))
}
