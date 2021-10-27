package main

import (
	"net/http"
	"strconv"

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
	id := r.URL.Query().Get("id")
	if id == "" {
		//send all the watch req
		i, err := app.store.GetAllUptimeWatchRequest(r.Context())
		if err != nil {
			sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
			return
		}
		sendResponse(rw, http.StatusOK, i, "")
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "invalid id")
		return
	}
	i, err := app.store.GetUptimeWatchRequestByID(r.Context(), intId)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		return
	}
	sendResponse(rw, http.StatusOK, i, "")
}
func (app *App) handleDeleteWatchReq(rw http.ResponseWriter, r *http.Request) {

}
