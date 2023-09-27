package controller

import (
	"net/http"
	"zsocket_service/model"
	"zsocket_service/response"
	"zsocket_service/validators"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) Login(ctx *gin.Context) {
	var (
		req model.Login
		err error
	)

	if err := validators.Binder(ctx, &req, c.log); err.Err != nil {
		err.Responder(ctx, c.log)

		return
	}

	if err := validators.Validator(ctx, &req, c.validate); err.Err != nil {
		err.Responder(ctx, c.log)

		return
	}

	loginResponse, err := c.service.Login(ctx, req)
	if err != nil {
		if responder := response.InternalError(ctx, err, c.log); responder.Err != nil {
			responder.Responder(ctx, c.log)

			return
		}
	}

	c.log.Info("loginResponse", zap.Any("data", loginResponse))

	ctx.JSON(http.StatusOK, loginResponse)
}
