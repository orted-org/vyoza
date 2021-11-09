package main

import "github.com/go-chi/chi/v5"

func initHandler(app *App, r *chi.Mux) {
	r.Post("/uptime", app.handleCreateWatchReq)
	r.Get("/uptime", app.handleGetWatchReq)
	r.Put("/uptime/{id}", app.handleUpdateWatchReq)
	r.Delete("/uptime/{id}", app.handleDeleteWatchReq)

	// config store
	r.Get("/cs/{name}", app.handleGetConfig)
	r.Post("/cs", app.handleSetConfig)
	r.Delete("/cs/{name}", app.handleDeleteConfig)
}
