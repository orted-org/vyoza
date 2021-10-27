package main

import "github.com/go-chi/chi/v5"

func initHandler(app *App, r *chi.Mux) {
	r.Post("/uptime", app.handleCreateWatchReq)
	r.Get("/uptime/{id}", app.handleGetWatchReq)
	r.Put("/uptime/{id}", app.handleUpdateWatchReq)
	r.Delete("/uptime/{id}", app.handleDeleteWatchReq)
}
