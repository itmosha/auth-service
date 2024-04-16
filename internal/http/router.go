package http

import (
	"net/http"

	"github.com/itmosha/auth-service/internal/controller"
)

// @Title Check service health
// @Success 200 {}
// @Resource Health
// @Route /health [get]
func NewRouter(controller *controller.Controller) (router *http.ServeMux) {
	router = http.NewServeMux()

	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	return
}
