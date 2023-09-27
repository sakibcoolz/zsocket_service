package validators

import (
	"net/http"
	"zsocket_service/errors"
	"zsocket_service/literals"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func Binder(ctx *gin.Context, obj any, log *zap.Logger) errors.ErrResp {
	var errs errors.ErrResp
	if err := ctx.ShouldBindBodyWith(obj, binding.JSON); err != nil {
		return ValidationError(ctx, err, log)
	}

	return errs
}

func Validator(ctx *gin.Context, obj any, validate *validator.Validate) errors.ErrResp {
	var errs errors.ErrResp

	if err := validate.Struct(obj); err != nil {
		errs = errors.ErrResp{
			Status: http.StatusBadRequest,
			Msg:    literals.ErrorResponse[literals.REQUEST_VALIDATION_ERROR],
			Err:    err,
		}
	}

	return errs
}

func ValidationError(ctx *gin.Context, err error, log *zap.Logger) errors.ErrResp {
	errs := errors.ErrResp{
		Status: http.StatusBadRequest,
		Msg:    literals.ErrorResponse[literals.REQUEST_VALIDATION_ERROR],
		Err:    err,
	}

	log.Error("Error in Body Validation", zap.Error(err), zap.Any("ctx", ctx), zap.Any("response", errs))

	return errs
}

func ValidationUnprocessableEntity(ctx *gin.Context, err error, log *zap.Logger) errors.ErrResp {
	errs := errors.ErrResp{
		Status: http.StatusUnprocessableEntity,
		Msg:    literals.ErrorResponse[literals.REQUEST_VALIDATION_ERROR],
		Err:    err,
	}

	log.Error("Error in Body Validation", zap.Error(err), zap.Any("ctx", ctx), zap.Any("response", errs))

	return errs
}
