package http

import "net/http"

func NewRouter() (router *http.ServeMux) {
	router = http.NewServeMux()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	return
}
