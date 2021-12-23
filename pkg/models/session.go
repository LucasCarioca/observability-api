package models

//Session model for tracking activity to a given session
type Session struct {
	DeviceId    string `json:"device_id" binding:"required"`
	AppId       string `json:"app_id" binding:"required"`
	AppVersion  string `json:"app_version" binding:"required"`
	BuildNumber string `json:"build_number" binding:"required"`
}

//SessionModel model for gorm
type SessionModel struct {
	Base
	Session
}
