package main

import "github.com/go-chi/chi/v5"

func initHandler(app *App, r *chi.Mux) {
	r.Post("/uptime", app.handleCreateWatchReq)
	r.Get("/uptime", app.handleGetWatchReq)
	r.Put("/uptime/{id}", app.handleUpdateWatchReq)
	r.Delete("/uptime/{id}", app.handleDeleteWatchReq)


	//Vault
	r.Get("/vault/{name}", app.handleGetVault)
	r.Post("/vault", app.handleSetVault)
	r.Put("/vault", app.handleUpdateVault)
	r.Delete("/vault/{name}", app.handleDeleteVault)
}
