package domain

import (
	"errors"
	"zsocket_service/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Store) GetDeviceInfo(ctx *gin.Context, login model.Login) (model.UserMaster, error) {
	var (
		device model.DeviceInfo
		user   model.UserMaster
	)

	if err := s.db.Model(&model.DeviceInfo{}).
		Where("device_key = ? and hostname = ?", login.DeviceKey, login.Hostname).
		Find(&device).Error; err != nil {
		s.log.Error("Failed to fetch DeviceInfo", zap.Error(err))

		return user, err
	}

	if device.DeviceKey == "" {
		err := errors.New("device_key not found")
		s.log.Error("Failed to fetch DeviceInfo", zap.Error(err))

		return user, err
	}

	if err := s.db.Model(&model.UserMaster{}).
		Where("tenant_master_id = ? and username = ?", device.UserMasterID, login.Username).
		Find(&user).Error; err != nil {
		s.log.Error("Failed to fetch DeviceInfo", zap.Error(err))

		return user, err
	}

	user.DeviceInfos = append(user.DeviceInfos, device)
	user.Password = ""

	return user, nil
}

func (s *Store) InsertSessionMaster(ctx *gin.Context, session model.SessionMaster) error {
	err := s.db.Save(&session).Error
	if err != nil {
		s.log.Error("error on insert", zap.Error(err))
		errs := s.db.Model(&model.SessionMaster{}).
			Where("device_key = ? and username = ? and hostname = ? and platform = ?",
				session.DeviceKey, session.Username, session.Hostname, session.Platform).
			Updates(map[string]interface{}{
				"session_id": session.SessionID,
				"status":     session.Status,
				"client_id":  session.ClientID,
			}).Error
		if errs != nil {
			s.log.Error("error ", zap.Error(err))

			return errs
		}

		err = errs
	}

	return err
}
