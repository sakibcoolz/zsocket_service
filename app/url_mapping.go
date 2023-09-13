package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"zsocket_service/config"
	"zsocket_service/controller"
	"zsocket_service/database"
	"zsocket_service/domain"
	"zsocket_service/event"
	"zsocket_service/eventhandler"
	"zsocket_service/literals"
	"zsocket_service/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

var (
	StopService = make(chan os.Signal, 1)
)

func Apps(config *config.Config, logger *zap.Logger, router *gin.Engine) *gin.Engine {
	go TerminateService(StopService, logger)

	signal.Notify(StopService, syscall.SIGINT, syscall.SIGTERM)

	validate := validator.New(validator.WithRequiredStructEnabled())

	db, err := database.Connection(config.GetCredentials())
	if err != nil {
		logger.Fatal("Failed to connect DB")
	}

	if err := database.Migrations(db.DB); err != nil {
		logger.Fatal("Failed in DB Migrations")
	}

	client := event.GetMqttConfig(config.GetMqtt())

	store := domain.NewStore(logger, db.DB)

	service := service.NewService(logger, store, client, config.GetMqtt().Topic)

	controller := controller.NewController(logger, service, validate)

	event := eventhandler.NewEvent(service, logger, client, config.GetMqtt().Topic)

	go event.Receiver()

	router.Use(gin.Recovery())

	preapproute := router.Group(fmt.Sprintf("/%s", literals.SERVICE))

	preapproute.POST("/login", controller.Login)

	return router
}

func TerminateService(stopService chan os.Signal, log *zap.Logger) {
	log.Info("Service Started")
	select {
	case <-stopService:
		log.Info("Terminating service")

		os.Exit(0)
	}
}
