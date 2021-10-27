package main

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
			sendErrorResponse(rw, http.StatusInternalServerError, nil, "internal server error")
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
		if err == sql.ErrNoRows {
			sendErrorResponse(rw, http.StatusNotFound, nil, "watch request with given id not found")
		} else {
			sendErrorResponse(rw, http.StatusInternalServerError, nil, "internal server error")
		}
		return
	}
	sendResponse(rw, http.StatusOK, i, "")
}
func (app *App) handleDeleteWatchReq(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "key not found in request")
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "invalid id")
		return
	}
	err = app.store.DeleteUptimeWatchRequestById(r.Context(), intId)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, "internal server error")
		return
	}
	sendResponse(rw, http.StatusOK, nil, "deleted watch request with id "+id)
}
