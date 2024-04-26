package controller

import "errors"

type (
	errorResponseBody struct {
		Message string `json:"message" example:"error description"`
	}

	CtxStatusCodeKey struct{}
	CtxErrorKey      struct{}
)

var (
	ErrServerError    = errors.New("internal server error")
	ErrNotImplemented = errors.New("not implemented")
)
