package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

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

	// auth service
	r.Post("/auth", app.handleLogin)
	r.Delete("/auth", app.handleCheckAllowance(http.HandlerFunc(app.handleLogout)))
}
