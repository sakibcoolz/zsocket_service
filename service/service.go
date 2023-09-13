package service

import (
	"zsocket_service/domain"
	"zsocket_service/model"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Service struct {
	log       *zap.Logger
	client    mqtt.Client
	store     domain.IStore
	mqtttopic string
}

type IService interface {
	Login(ctx *gin.Context, login model.Login) (model.LoginResponse, error)
	MsgFiner(msg model.Msg)
}

func NewService(logger *zap.Logger, store domain.IStore, client mqtt.Client, topic string) *Service {
	return &Service{
		log:       logger,
		store:     store,
		client:    client,
		mqtttopic: topic,
	}
}
