package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initHandler(app *App, r *chi.Mux) {

	// uptime
	r.Post("/uptime", app.handleCheckAllowance(http.HandlerFunc(app.handleCreateWatchReq)))
	r.Get("/uptime", app.handleCheckAllowance(http.HandlerFunc(app.handleGetWatchReq)))
	r.Put("/uptime/{id}", app.handleCheckAllowance(http.HandlerFunc(app.handleUpdateWatchReq)))
	r.Delete("/uptime/{id}", app.handleCheckAllowance(http.HandlerFunc(app.handleDeleteWatchReq)))

	// vault
	r.Get("/vault/{name}", app.handleGetVault)
	r.Post("/vault", app.handleCheckAllowance(http.HandlerFunc(app.handleSetVault)))
	r.Put("/vault", app.handleCheckAllowance(http.HandlerFunc(app.handleUpdateVault)))
	r.Delete("/vault/{name}", app.handleCheckAllowance(http.HandlerFunc(app.handleDeleteVault)))

	// config store
	r.Get("/cs/{name}", app.handleGetConfig)
	r.Post("/cs", app.handleSetConfig)
	r.Delete("/cs/{name}", app.handleDeleteConfig)

	// services
	r.Post("/service", app.handleCheckAllowance(http.HandlerFunc(app.handleCreateService)))
	r.Delete("/service/{id}", app.handleCheckAllowance(http.HandlerFunc(app.handleDeleteService)))
	r.Get("/service/{id}", app.handleCheckAllowance(http.HandlerFunc(app.handleGetService)))

	// auth service
	r.Post("/auth", app.handleLogin)
	r.Delete("/auth", app.handleCheckAllowance(http.HandlerFunc(app.handleLogout)))
}
