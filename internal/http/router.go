package http

import "net/http"

// @Title Check service health
// @Success 200 {}
// @Resource Health
// @Route /health [get]
func NewRouter() (router *http.ServeMux) {
	router = http.NewServeMux()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	return
}
