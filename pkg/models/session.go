package models

import (
	//"github.com/google/uuid"
	"gorm.io/gorm"
)

//Session model for tracking activity to a given session
type Session struct {
	DeviceID    string `json:"device_id" binding:"required"`
	AppID       string `json:"app_id" binding:"required"`
	AppVersion  string `json:"app_version" binding:"required"`
	BuildNumber string `json:"build_number" binding:"required"`
}

//SessionModel model for gorm
type SessionModel struct {
	Base
	Session
	SessionKey string `json:"session_key" binding:"required"`
}

//BeforeCreate creates a random uuid registration key for new invitations
func (s *SessionModel) BeforeCreate(tx *gorm.DB) error {
	s.SessionKey = "thisisaplaceholderwhileonline" //uuid.NewString()
	return nil
}
