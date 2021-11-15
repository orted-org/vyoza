package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	db "github.com/orted-org/vyoza/db/dao"
)

// TODO: input validation, register in watcher
func (app *App) handleSetConfig(rw http.ResponseWriter, r *http.Request) {
	var arg db.KeyValue
	err := getBody(r, &arg)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}

	// input validation
	err = validateSetConfig(r.Context(), arg)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}

	// create in store
	err = app.configStore.Set(r.Context(), arg.Key, arg.Value)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		return
	}

	sendResponse(rw, http.StatusCreated, arg, "config set for "+arg.Key)
}

// TODO: input validation
func (app *App) handleGetConfig(rw http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "config name not found in request")
		return
	}
	i, err := app.configStore.Get(r.Context(), name)
	if err != nil {
		if err == sql.ErrNoRows {
			sendErrorResponse(rw, http.StatusNotFound, nil, "config not found with name "+name)
		} else {
			sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		}
		return
	}
	sendResponse(rw, http.StatusOK, i, "")
}
func (app *App) handleDeleteConfig(rw http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "config name not found in request")
		return
	}
	err := app.configStore.Delete(r.Context(), name)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		return
	}
	sendResponse(rw, http.StatusOK, nil, "config deleted with name "+name)
}
func validateSetConfig(ctx context.Context, i db.KeyValue) error {
	return validation.ValidateStructWithContext(ctx, &i,

		validation.Field(&i.Key, validation.Required),
		validation.Field(&i.Value, validation.Required),
	)
}
