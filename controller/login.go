package controller

import (
	"net/http"
	"zsocket_service/model"
	handler "zsocket_service/request"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Login(ctx *gin.Context) {
	var (
		req model.Login
		err error
	)

	if err := handler.Binder(ctx, &req, c.log); err != nil {
		err.Responder(ctx)

		return
	}

	if err := handler.Validator(ctx, &req, c.validate); err != nil {
		err.Responder(ctx)

		return
	}

	loginResponse, err := c.service.Login(ctx, req)
	if err != nil {
		if responder := handler.InternalError(ctx, err, c.log); responder != nil {
			responder.Responder(ctx)

			return
		}
	}

	ctx.JSON(http.StatusOK, loginResponse)
}
