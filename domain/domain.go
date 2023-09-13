package domain

import (
	"zsocket_service/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Store struct {
	log *zap.Logger
	db  *gorm.DB
}

type IStore interface {
	GetDeviceInfo(ctx *gin.Context, login model.Login) (model.UserMaster, error)
	InsertSessionMaster(ctx *gin.Context, session model.SessionMaster) error
}

func NewStore(logger *zap.Logger, db *gorm.DB) *Store {
	return &Store{
		log: logger,
		db:  db,
	}
}
