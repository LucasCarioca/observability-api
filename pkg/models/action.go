package models

//Action model for actions
type Action struct {
	Action  string `json:"action" binding:"required"`
	Details string `json:"details"`
}

//ActionModel gorm model for actions
type ActionModel struct {
	Base
	Action
	SessionID uint         `json:"session_id"`
	Session   SessionModel `json:"session"`
}
