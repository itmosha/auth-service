package controller

import (
	"context"
	"encoding/json"
	"net/http"
)

type (
	SuccessResponseBody interface{}
	ErrorResponseBody   struct {
		Message string `json:"message" example:"error description"`
	}

	CtxStatusCodeKey struct{}
	CtxErrorKey      struct{}
)

func readBodyToStruct[T any](r *http.Request, out *T) (*T, error) {
	err := json.NewDecoder(r.Body).Decode(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func ResponseWithSuccess(w http.ResponseWriter, r *http.Request, statusCode int, body interface{}) {
	w.WriteHeader(statusCode)
	if body != nil {
		json.NewEncoder(w).Encode(body)
	}
	ctx := r.Context()
	ctx = context.WithValue(ctx, CtxStatusCodeKey{}, statusCode)
	req := r.WithContext(ctx)
	*r = *req
}

func ResponseWithError(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	w.WriteHeader(statusCode)
	if err != nil {
		json.NewEncoder(w).Encode(ErrorResponseBody{Message: err.Error()})
	}
	ctx := r.Context()
	ctx = context.WithValue(ctx, CtxStatusCodeKey{}, statusCode)
	ctx = context.WithValue(ctx, CtxErrorKey{}, err)
	req := r.WithContext(ctx)
	*r = *req
}
