package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	HTTP "net/http"

	"github.com/itmosha/auth-service/internal/config"
	"github.com/itmosha/auth-service/internal/controller"
	"github.com/itmosha/auth-service/internal/http"
	"github.com/itmosha/auth-service/internal/http/middleware"
	storagePostgres "github.com/itmosha/auth-service/internal/storage/postgres"
	storageRedis "github.com/itmosha/auth-service/internal/storage/redis"
	"github.com/itmosha/auth-service/internal/usecase"
	"github.com/itmosha/auth-service/pkg/clients/postgres"
	"github.com/itmosha/auth-service/pkg/clients/redis"
	"github.com/itmosha/auth-service/pkg/logger"
	"github.com/itmosha/simplejwt"
)

func Run(cfg *config.Config) {
	pgClient, err := postgres.NewPostgresClient(&cfg.DB)
	if err != nil {
		log.Fatalf("could not create postgres client: %v\n", err)
	}
	redisClient := redis.NewRedisClient(&cfg.Cache)
	logger := logger.NewLogger("logs/"+cfg.HTTPServer.LogFileName, cfg.Env)
	jwtClient := simplejwt.NewJWTClient([]byte(cfg.JWTSecret), nil)

	storage := storagePostgres.NewStoragePostgres(pgClient)
	cache := storageRedis.NewCacheRedis(redisClient)

	go func() {
		for {
			ctx := context.Background()
			err := storage.DeleteUnregistered(&ctx, time.Minute*30)
			if err != nil {
				logger.LogError("storagePostgres.DeleteUnregistered", err)
			}
			time.Sleep(time.Minute)
		}
	}()

	usecase := usecase.NewUsecase(storage, cache, jwtClient)
	controller := controller.NewController(usecase, logger)

	router := http.NewRouter(controller)

	server := &HTTP.Server{
		Handler:      middleware.LoggerMiddleware(logger, router),
		Addr:         fmt.Sprintf(":%s", cfg.HTTPServer.RunPort),
		WriteTimeout: cfg.HTTPServer.Timeout,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	serverErrs := make(chan error, 1)
	go func() {
		log.Println("Server starting on port", cfg.HTTPServer.RunPort)
		serverErrs <- server.ListenAndServe()
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	shutdown := gracefulShutdown(server)

	select {
	case err := <-serverErrs:
		shutdown(err)
	case sig := <-quit:
		shutdown(sig)
	}
	log.Println("Server exiting")
}

func gracefulShutdown(server *HTTP.Server) func(reason interface{}) {
	return func(reason interface{}) {
		log.Println("Service graceful shutdown:", reason)

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Println("Service graceful shutdown Failed:", err)
		}
	}
}
