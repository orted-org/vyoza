package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	is "github.com/go-ozzo/ozzo-validation/v4/is"
	db "github.com/orted-org/vyoza/db/dao"
	"github.com/orted-org/vyoza/internal/watcher"
	"github.com/orted-org/vyoza/util"
)

// TODO: input validation, register in watcher
func (app *App) handleCreateWatchReq(rw http.ResponseWriter, r *http.Request) {
	var arg db.AddUptimeWatchRequestParams
	err := getBody(r, &arg)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}

	// input validation
	err = validateCreateWatchRequest(r.Context(), arg)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}

	// sha256 the hook_secret
	if arg.HookSecret != "" {
		arg.HookSecret = string(util.NewSHA256([]byte(arg.HookSecret)))
	}

	// create in store
	i, err := app.store.AddUptimeWatchRequest(r.Context(), arg)
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		return
	}

	_, err = app.store.AddUptimeSSLInfo(r.Context(), db.UptimeSSLInfo{
		UWRID: i.ID,
	})
	if err != nil {
		app.store.DeleteUptimeWatchRequestById(r.Context(), i.ID)
		sendErrorResponse(rw, http.StatusInternalServerError, nil, err.Error())
		return
	}

	// registering the watcher if enabled
	if i.Enabled {
		app.watcher.Register(watcher.WatcherParams{
			ID:              i.ID,
			Location:        i.Location,
			Interval:        i.Interval,
			ExpectedStatus:  i.ExpectedStatus,
			MaxResponseTime: i.MaxResponseTime,
		})
	}
	if i.SSLMonitor {
		app.watcher.RegisterSSL(watcher.SSLWatcherParams{
			ID:       i.ID,
			Location: i.Location,
			Interval: i.SSLInterval,
		})
	}

	sendResponse(rw, http.StatusCreated, i, "created uptime watch request")
}

// TODO: input validation
func (app *App) handleUpdateWatchReq(rw http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "key not found in request")
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, "invalid id")
		return
	}

	var arg map[string]interface{}

	// getting body
	err = getBody(r, &arg)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}

	for k, v := range arg {
		switch v.(type) {
		case float32, float64, int16, int32, int64:
			if d, ok := v.(float64); ok {
				arg[k] = int(d)
			} else {
				sendErrorResponse(rw, http.StatusBadRequest, nil, "unsupported field type")
				return
			}
		}
	}

	// input validation
	err = validateUpdateWatchRequest(r.Context(), arg)
	if err != nil {
		sendErrorResponse(rw, http.StatusBadRequest, nil, err.Error())
		return
	}

	// sha256 the hook_secret
	if secret, ok := arg["hook_secret"].(string); ok {
		if secret != "" {
			secret = string(util.NewSHA256([]byte(secret)))
			arg["hook_secret"] = secret
		}
	}

	// update in store
	i, err := app.store.UpdateUptimeWatchRequestById(r.Context(), arg, intId)
	if err != nil {
		if err == sql.ErrNoRows {
			sendErrorResponse(rw, http.StatusNotFound, nil, "watch request with given id not found")
		} else {
			sendErrorResponse(rw, http.StatusInternalServerError, nil, "internal server error")
		}
		return
	}

	// registering the watcher if enabled
	if i.Enabled {
		app.watcher.Register(watcher.WatcherParams{
			ID:              i.ID,
			Location:        i.Location,
			Interval:        i.Interval,
			ExpectedStatus:  i.ExpectedStatus,
			MaxResponseTime: i.MaxResponseTime,
		})
	}
	if i.SSLMonitor {
		app.watcher.RegisterSSL(watcher.SSLWatcherParams{
			ID:       i.ID,
			Location: i.Location,
			Interval: i.SSLInterval,
		})
	}
	sendResponse(rw, http.StatusOK, i, "")
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
		return
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
	app.store.DeleteUptimeSSLInfoByUWRID(r.Context(), intId)

	// un registering the watcher
	app.watcher.UnRegsiter(intId)
	app.watcher.UnRegsiterSSL(intId)
	sendResponse(rw, http.StatusOK, nil, "deleted watch request with id "+id)
}

func validateCreateWatchRequest(ctx context.Context, i db.AddUptimeWatchRequestParams) error {
	return validation.ValidateStructWithContext(ctx, &i,

		validation.Field(&i.Title, validation.Required),
		validation.Field(&i.Location, validation.Required, is.URL),

		// interval is in seconds and hence min interval is 20s
		validation.Field(&i.Interval, validation.Required, validation.Min(20)),
		validation.Field(&i.SSLInterval, validation.Min(20)),

		// exp notification is in hrs and hence minimum is 24hrs
		validation.Field(&i.SSLExpiryNotification, validation.Min(24)),

		validation.Field(&i.ExpectedStatus, validation.Required, validation.Min(100), validation.Max(599)),

		// response time is in milliseconds
		validation.Field(&i.StdResponseTime, validation.Required, validation.Min(1)),
		validation.Field(&i.MaxResponseTime, validation.Required, validation.Min(i.StdResponseTime)),

		// retain duration is in hrs and hence minimum is 24hrs
		validation.Field(&i.RetainDuration, validation.Required, validation.Min(24)),

		validation.Field(&i.HookAddress, is.URL),
		validation.Field(&i.HookSecret, validation.Length(16, 0)),
		validation.Field(&i.NotificationEmail, is.Email),
	)
}

func validateUpdateWatchRequest(ctx context.Context, arg map[string]interface{}) error {

	var i db.AddUptimeWatchRequestParams
	js, err := json.Marshal(arg)
	if err != nil {
		return ErrCouldNotParseBody
	}
	err = json.Unmarshal(js, &i)
	if err != nil {
		return ErrCouldNotParseBody
	}

	return validation.ValidateStructWithContext(ctx, &i,

		validation.Field(&i.Title),
		validation.Field(&i.Location, is.URL),

		// interval is in seconds and hence min interval is 20s
		validation.Field(&i.Interval, validation.Min(20)),
		validation.Field(&i.SSLInterval, validation.Min(20)),

		// exp notification is in hrs and hence minimum is 24hrs
		validation.Field(&i.SSLExpiryNotification, validation.Min(24)),

		validation.Field(&i.ExpectedStatus, validation.Min(100), validation.Max(599)),

		// response time is in milliseconds
		validation.Field(&i.StdResponseTime, validation.Min(1)),
		validation.Field(&i.MaxResponseTime, validation.Min(i.StdResponseTime)),

		// retain duration is in hrs and hence minimum is 24hrs
		validation.Field(&i.RetainDuration, validation.Min(24)),

		validation.Field(&i.HookAddress, is.URL),
		validation.Field(&i.HookSecret, validation.Length(16, 0)),
		validation.Field(&i.NotificationEmail, is.Email),
	)
}
