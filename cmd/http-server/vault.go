package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	db "github.com/orted-org/vyoza/db/dao"
)

// TODO:register in watcher
func (app *App) handleSetVault(rw http.ResponseWriter, r *http.Request) {
	var arg db.KeyValue
	err := getBody(r, &arg)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}

	// input validation
	err = validateVaultInp(r.Context(), arg)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}

	// create in store
	err = app.vault.Set(r.Context(), arg.Key, arg.Value)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		return
	}

	sendResponse(rw, http.StatusCreated, arg, "Vault set for "+arg.Key)
}

func (app *App) handleUpdateVault(rw http.ResponseWriter, r *http.Request) {
	var arg db.KeyValue
	err := getBody(r, &arg)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}

	// input validation
	err = validateVaultInp(r.Context(), arg)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}

	// create in store
	err = app.vault.Update(r.Context(), arg.Key, arg.Value)
	if err != nil {
		if err == sql.ErrNoRows {
			sendErrorResponse(rw, http.StatusNotFound, nil, "vault not found with name "+arg.Key)
		} else {
			sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		}
		return
	}

	sendResponse(rw, http.StatusCreated, arg, "Vault updated for "+arg.Key)
}

func (app *App) handleGetVault(rw http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "vault name not found in request")
		return
	}
	i, err := app.vault.Get(r.Context(), name)
	if err != nil {
		if err == sql.ErrNoRows {
			sendErrorResponse(rw, http.StatusNotFound, nil, "vault not found with name "+name)
		} else {
			sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		}
		return
	}

	sendResponse(rw, http.StatusOK, i, "")
}
func (app *App) handleDeleteVault(rw http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "vault name not found in request")
		return
	}
	err := app.vault.Delete(r.Context(), name)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		return
	}
	sendResponse(rw, http.StatusOK, nil, "vault deleted with name "+name)
}
func validateVaultInp(ctx context.Context, i db.KeyValue) error {
	return validation.ValidateStructWithContext(ctx, &i,

		validation.Field(&i.Key, validation.Required),
		validation.Field(&i.Value, validation.Required),
	)
}
