package controller

import (
	"zsocket_service/service"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Controller struct {
	log      *zap.Logger
	service  service.IService
	validate *validator.Validate
}

func NewController(log *zap.Logger, service service.IService, validate *validator.Validate) *Controller {
	return &Controller{
		log:      log,
		service:  service,
		validate: validate,
	}
}
