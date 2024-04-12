package main

import (
	"github.com/itmosha/auth-service/internal/config"
	"github.com/itmosha/auth-service/internal/service"
)

func main() {
	cfg := config.NewConfig()
	service.Run(cfg)
}
