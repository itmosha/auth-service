package controller

type (
	successResponseBody interface{}
	errorResponseBody   struct {
		Message string `json:"message" example:"error description"`
	}

	CtxStatusCodeKey struct{}
	CtxErrorKey      struct{}
)
