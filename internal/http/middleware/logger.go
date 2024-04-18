package middleware

import (
	"net/http"

	"github.com/itmosha/auth-service/internal/controller"
	"github.com/itmosha/auth-service/pkg/logger"
)

func LoggerMiddleware(logger *logger.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		var statusCode int
		var err error

		statusCode, _ = r.Context().Value(controller.CtxStatusCodeKey{}).(int)
		if statusCode >= 400 {
			err, _ = r.Context().Value(controller.CtxErrorKey{}).(error)
		}
		logger.Log(r, statusCode, err)
	})
}
