package service

import (
	"fmt"
	"log"

	HTTP "net/http"

	"github.com/itmosha/auth-service/internal/config"
	"github.com/itmosha/auth-service/internal/http"
	"github.com/itmosha/auth-service/pkg/logger"
	"github.com/itmosha/auth-service/pkg/postgres"
)

func Run(cfg *config.Config) {
	pgClient, err := postgres.NewPostgresClient(&cfg.DB)
	if err != nil {
		log.Fatalf("could not create postgres client: %v\n", err)
	}
	logger := logger.NewLogger("logs/auth.log", cfg.Env)

	_ = pgClient
	_ = logger

	router := http.NewRouter()

	server := &HTTP.Server{
		Handler:      router,
		Addr:         fmt.Sprintf(":%s", cfg.HTTPServer.RunPort),
		WriteTimeout: cfg.HTTPServer.Timeout,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}
	log.Printf("starting server on port %s\n", cfg.HTTPServer.RunPort)
	log.Fatal(server.ListenAndServe())
}
