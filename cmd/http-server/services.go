package main

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	db "github.com/orted-org/vyoza/db/dao"
	"github.com/orted-org/vyoza/util"
)

func (app *App) handleCreateService(rw http.ResponseWriter, r *http.Request) {
	var arg db.Service
	err := getBody(r, &arg)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}

	if arg.Name == "" {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "name not provided")
		return
	}

	id := strings.ReplaceAll(uuid.New().String(), "-", "")
	secret := util.RandomAlphaNumericSymbolString(64)
	hash, err := util.HashSecret(secret)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		return
	}

	// replacing the incoming(if any) id and hash with auto generated one
	arg.ID = id
	arg.Secret = hash

	_, err = app.store.AddService(r.Context(), arg)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		return
	}

	res := make(map[string]interface{})
	res["API-KEY"] = id
	res["API-SECRET"] = secret
	sendResponse(rw, http.StatusCreated, res, "created service")
}

func (app *App) handleDeleteService(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "config name not found in request")
		return
	}
	err := app.store.DeleteServiceByID(r.Context(), id)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		return
	}
	sendResponse(rw, http.StatusCreated, nil, "service deleted with id "+id)
}

func (app *App) handleGetService(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		// send all the services
		items, err := app.store.GetAllService(r.Context())
		if err != nil {
			sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
			return
		}
		sendResponse(rw, http.StatusCreated, items, "")
		return
	}
	item, err := app.store.GetServiceByID(r.Context(), id)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		return
	}
	sendResponse(rw, http.StatusCreated, item, "")
}
