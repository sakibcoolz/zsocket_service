package response

import (
	"net/http"
	"zsocket_service/errors"
	"zsocket_service/literals"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func InternalError(ctx *gin.Context, err error, log *zap.Logger) errors.ErrResp {
	errs := errors.ErrResp{
		Status: http.StatusInternalServerError,
		Msg:    literals.ErrorResponse[literals.FAIL_TO_STORE],
		Err:    err,
	}

	log.Error("Error in store", zap.Error(err), zap.Any("ctx", ctx), zap.Any("response", errs))

	return errs
}

func LoginFailed(ctx *gin.Context, err error, log *zap.Logger) errors.ErrResp {
	errs := errors.ErrResp{
		Status: http.StatusNoContent,
		Msg:    literals.ErrorResponse[literals.LOGIN_FAILED],
		Err:    err,
	}

	log.Error("Error in Body Validation", zap.Error(err), zap.Any("ctx", ctx), zap.Any("response", errs))

	return errs
}
