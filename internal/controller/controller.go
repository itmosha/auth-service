package controller

import (
	"encoding/json"
	"net/http"
)

type SuccessResponseBody interface{}

func ResponseWithSuccess(w http.ResponseWriter, statusCode int, body interface{}) {
	w.WriteHeader(statusCode)
	if body != nil {
		json.NewEncoder(w).Encode(body)
	}
}

type ErrorResponseBody struct {
	Message string `json:"message"`
}

func ResponseWithError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponseBody{Message: err.Error()})
	}
}
