package main

import (
	"fmt"
	"zsocket_service/app"
	"zsocket_service/config"
	"zsocket_service/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger := log.NewLogger()

	config := config.NewConfig(logger)

	router := gin.Default()

	router = app.Apps(config, logger, router)

	if err := router.Run(fmt.Sprintf(":%d", config.GetServiceConfig().Port)); err != nil {
		logger.Fatal("Service Error", zap.Error(err))
	}
}
