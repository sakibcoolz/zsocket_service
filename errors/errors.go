package errors

import (
	"encoding/json"
	"zsocket_service/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ErrResp struct {
	Status int    `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Err    error
}

type IErrResp interface {
	Responder(ctx *gin.Context)
}

func (e ErrResp) Error() string {
	return e.String()
}

func (e ErrResp) String() string {
	msg, _ := json.Marshal(model.ErrFormat{
		Status: e.Status,
		Msg:    e.Msg,
		Err:    e.Err.Error(),
	})

	return string(msg)
}

func (e ErrResp) Responder(ctx *gin.Context, log *zap.Logger) {
	log.Error("responding with error", zap.Error(e))
	ctx.JSON(e.Status, e.Msg)
}
