package controller

import (
	"github.com/itmosha/auth-service/pkg/logger"
)

type Controller struct {
	logger *logger.Logger
}

func NewController(logger *logger.Logger) *Controller {
	return &Controller{logger}
}
