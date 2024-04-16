package main

import (
	"github.com/itmosha/auth-service/internal/config"
	"github.com/itmosha/auth-service/internal/service"
)

// @Version 0.0.1
// @Title Auth Service
// @Description Auth service which is responsible for user authentication and authorization.
// @Security Authorization
// @SecurityScheme JWT http bearer Your JWT token
func main() {
	cfg := config.NewConfig()
	service.Run(cfg)
}
