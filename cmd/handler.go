package main

import "github.com/go-chi/chi/v5"

func initHandler(app *App, r *chi.Mux) {
	r.Post("/uptime", app.handleCreateWatchReq)
	r.Get("/uptime", app.handleGetWatchReq)
	r.Put("/uptime", app.handleUpdateWatchReq)
	r.Delete("/uptime", app.handleDeleteWatchReq)
}
