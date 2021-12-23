package models

//Action model for actions
type Action struct {
	Action    string       `json:"action" binding:"required"`
	SessionId uint         `json:"session_id" binding:"required"`
	Session   SessionModel `json:"session" binding:"required"`
}

//ActionModel gorm model for actions
type ActionModel struct {
	Base
}
