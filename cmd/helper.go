package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var (
	ErrCouldNotReadBody  = errors.New("could not read body")
	ErrCouldNotParseBody = errors.New("could not parse body")
)

func getBody(r *http.Request, v interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return ErrCouldNotReadBody
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		return ErrCouldNotReadBody
	}
	return nil
}

type httpResp struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func sendResponse(rw http.ResponseWriter, status int, data interface{}, message string) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(status)
	out, err := json.Marshal(httpResp{Status: "success", Data: data, Message: message})
	if err != nil {
		sendErrorResponse(rw, http.StatusInternalServerError, nil, "Internal Server Error")
		return
	}

	rw.Write(out)
}
func sendErrorResponse(rw http.ResponseWriter, status int, data interface{}, message string) {
	rw.Header().Set("Content-Type", "application/json; charset=utf-8")
	rw.WriteHeader(status)
	out, _ := json.Marshal(httpResp{Status: "error",
		Message: message,
		Data:    data})

	rw.Write(out)
}
