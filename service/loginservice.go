package service

import (
	"errors"
	"fmt"
	"zsocket_service/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Service) Login(ctx *gin.Context, login model.Login) (model.LoginResponse, error) {
	var (
		sessionmaster model.SessionMaster
		loginResponse model.LoginResponse
		err           error
	)

	userinfo, err := s.store.GetDeviceInfo(ctx, login)
	if err != nil {
		return loginResponse, err
	}

	if len(userinfo.DeviceInfos) == 0 {
		return loginResponse, errors.New("No device found")
	}

	sessionid := uuid.New().String()
	clientid := fmt.Sprintf("%s_%s", "FIRE_CLIENT", sessionid)

	sessionmaster.ClientID = clientid
	sessionmaster.SessionID = sessionid
	sessionmaster.Hostname = userinfo.DeviceInfos[0].Hostname
	sessionmaster.Topic = s.mqtttopic
	sessionmaster.Platform = userinfo.DeviceInfos[0].Platform
	sessionmaster.DeviceKey = userinfo.DeviceInfos[0].DeviceKey
	sessionmaster.Username = login.Username
	sessionmaster.Status = "ACTIVE"

	if err := s.store.InsertSessionMaster(ctx, sessionmaster); err != nil {
		return loginResponse, err
	}

	loginResponse = model.LoginResponse{
		ClientID:  clientid,
		SessionID: sessionid,
		Topic:     s.mqtttopic,
	}

	return loginResponse, nil
}
