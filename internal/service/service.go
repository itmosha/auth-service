package service

import (
	"fmt"
	"log"

	HTTP "net/http"

	"github.com/itmosha/auth-service/internal/config"
	"github.com/itmosha/auth-service/internal/controller"
	"github.com/itmosha/auth-service/internal/http"
	storage "github.com/itmosha/auth-service/internal/storage/postgres"
	"github.com/itmosha/auth-service/internal/usecase"
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

	storage := storage.NewAuthStoragePostgres(pgClient)
	usecase := usecase.NewAuthUsecase(storage)
	controller := controller.NewController(usecase, logger)

	router := http.NewRouter(controller)

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
