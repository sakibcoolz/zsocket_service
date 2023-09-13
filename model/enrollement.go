package model

import (
	"time"

	"gorm.io/gorm"
)

type DeviceInfo struct {
	gorm.Model
	UserMasterID uint      `gorm:"column:user_master_id" json:"user_master_id,omitempty"`
	Hostname     string    `gorm:"column:hostname" json:"hostname,omitempty"`
	DeviceKey    string    `gorm:"column:device_key;index:idx_keys,unique,priority:1" json:"device_key,omitempty"`
	Platform     string    `gorm:"platform" json:"platform,omitempty"`
	OS           string    `gorm:"os" json:"os,omitempty"`
	Status       string    `gorm:"status" json:"status,omitempty"`
	ExpiryDate   time.Time `gorm:"expiry_date" json:"expiry_date,omitempty"`
}

func (DeviceInfo) TableName() string {
	return "zregistry.device_infos"
}

// User AS in User / Device
type UserMaster struct {
	gorm.Model
	TenantMasterID uint         `gorm:"column:tenant_master_id;index:idx_member,unique,priority:1" json:"tenant_master_id,omitempty"`
	Name           string       `gorm:"column:name" json:"name,omitempty"`
	Username       string       `gorm:"column:username;index:idx_member,unique,priority:2" json:"username,omitempty"`
	ExpiryDate     time.Time    `gorm:"column:expiry_date" json:"expiry_date,omitempty"`
	UserType       string       `gorm:"column:user_type" json:"user_type,omitempty"`
	Password       string       `gorm:"column:password" json:"password,omitempty"`
	Address        string       `gorm:"column:address" json:"address,omitempty"`
	DeviceInfos    []DeviceInfo `json:"device_infos,omitempty"`
}

func (UserMaster) TableName() string {
	return "zregistry.user_masters"
}

type AgentConfig struct {
	gorm.Model
	EnrollmentUrl string `gorm:"column:enrollment_url" json:"enrollmentUrl"`
	EnrolledUrl   string `gorm:"column:enrolled_url" json:"enrolledUrl"`
	MQTT          string `gorm:"column:mqtt" json:"mqtt"`
	Topic         string `gorm:"column:topic" json:"topic"`
}

type TopicMaster struct {
	gorm.Model
	Topic  string
	Status string
}

type SessionMaster struct {
	gorm.Model
	Hostname  string `gorm:"column:hostname" json:"hostname,omitempty"`
	SessionID string `gorm:"column:session_id" json:"sessionID,omitempty"`
	Status    string `gorm:"column:status"`
	Platform  string `gorm:"column:platform" json:"platform,omitempty"`
	ClientID  string `gorm:"column:client_id" json:"client_id,omitempty"`
	Username  string `gorm:"column:username" json:"username,omitempty"`
	Topic     string `gorm:"column:topic" json:"topic"`
	DeviceKey string `gorm:"column:device_key;index:idx_member,unique,priority:1" json:"device_key,omitempty"`
}

// Request body
type Login struct {
	Username  string `gorm:"username;index:idx_username,unique" json:"username,omitempty" validate:"required"`
	DeviceKey string `gorm:"column:device_key" json:"device_key,omitempty" validate:"required"`
	Hostname  string `gorm:"hostname" json:"hostname,omitempty" validate:"required"`
}

type LoginResponse struct {
	ClientID  string `gorm:"client_id" json:"client_id,omitempty"`
	SessionID string `gorm:"column:session_id" json:"sessionID,omitempty"`
	Topic     string `gorm:"column:topic" json:"topic"`
}
