package main

import (
	"net/http"

	db "github.com/orted-org/vyoza/db/dao"
)

func (app *App) handleCreateWatchReq(rw http.ResponseWriter, r *http.Request) {
	var arg db.AddUptimeWatchRequestParams
	err := getBody(r, &arg)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		return
	}
	i, err := app.store.AddUptimeWatchRequest(r.Context(), arg)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		return
	}
	sendResponse(rw, http.StatusCreated, i, "created uptime watch request")
}
func (app *App) handleUpdateWatchReq(rw http.ResponseWriter, r *http.Request) {

}
func (app *App) handleGetWatchReq(rw http.ResponseWriter, r *http.Request) {

}
func (app *App) handleDeleteWatchReq(rw http.ResponseWriter, r *http.Request) {

}
