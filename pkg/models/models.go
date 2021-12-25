package models

import (
	"gorm.io/gorm"
)

//Init initializes all models and creates tables if not already present
func Init(db *gorm.DB) {
	db.AutoMigrate(&SessionModel{})
	db.AutoMigrate(&ActionModel{})
}
