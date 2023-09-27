package model

import (
	"time"

	"gorm.io/gorm"
)

type Register struct {
	Tenent TenantMaster `json:"tenent,omitempty"`
	User   UserMaster   `json:"user,omitempty"`
}

type TenantMaster struct {
	gorm.Model
	Name       string       `gorm:"column:name" json:"name,omitempty"`
	Status     string       `gorm:"column:status" json:"status,omitempty"` // Active / block / Lock
	Domain     string       `gorm:"domain;index:idx_domain,unique" json:"domain,omitempty"`
	ExpiryDate time.Time    `gorm:"column:expiry_date" json:"expiry_date,omitempty"`
	Users      []UserMaster `json:"users,omitempty"`
}

type Contact struct {
	gorm.Model
	UserMasterID uint   `gorm:"column:user_master_id" json:"user_master_id,omitempty"`
	Types        string `gorm:"column:types" json:"types,omitempty"`
	Value        string `gorm:"column:value" json:"value,omitempty"`
	Status       string `gorm:"column:status" json:"status,omitempty"` // Active / Block / Lock
}

type ErrFormat struct {
	Status int    `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Err    string `json:"err,omitempty"`
}
