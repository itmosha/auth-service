package controller

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrDecodeBody  = errors.New("could not decode body")
	ErrServerError = errors.New("internal server error")
)

type SuccessResponseBody interface{}

type ErrorResponseBody struct {
	Message string `json:"message"`
}

func readBodyToStruct[T any](r *http.Request, out *T) (*T, error) {
	err := json.NewDecoder(r.Body).Decode(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func ResponseWithSuccess(w http.ResponseWriter, statusCode int, body interface{}) {
	w.WriteHeader(statusCode)
	if body != nil {
		json.NewEncoder(w).Encode(body)
	}
}

func ResponseWithError(w http.ResponseWriter, statusCode int, err error) {
	w.WriteHeader(statusCode)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponseBody{Message: err.Error()})
	}
}
