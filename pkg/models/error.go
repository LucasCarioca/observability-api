package models

type Error struct {
	Message string `json:"message" binding:"required"`
	Code    string `json:"code" binding:"required"`
}
