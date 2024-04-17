package http

import (
	"net/http"
)

type AuthHandlerInterface interface {
	Register() http.HandlerFunc
	ConfirmRegister() http.HandlerFunc
	Login() http.HandlerFunc
	ConfirmLogin() http.HandlerFunc
	Refresh() http.HandlerFunc
	Revoke() http.HandlerFunc
}

// @Title Check service health
// @Success 200 {}
// @Resource Health
// @Route /health [get]
func NewRouter(controller AuthHandlerInterface) (router *http.ServeMux) {
	router = http.NewServeMux()

	router.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("POST /api/register/", controller.Register())
	router.HandleFunc("POST /api/register/confirm/", controller.ConfirmRegister())
	router.HandleFunc("POST /api/login/", controller.Login())
	router.HandleFunc("POST /api/login/confirm/", controller.ConfirmLogin())
	router.HandleFunc("POST /api/refresh/", controller.Refresh())
	router.HandleFunc("POST /api/revoke/", controller.Revoke())

	return
}
