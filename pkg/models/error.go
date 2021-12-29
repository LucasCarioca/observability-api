package models

//Error internal error model to be provided to integrated systems
type Error struct {
	Message string `json:"message" binding:"required"`
	Code    string `json:"code" binding:"required"`
}
