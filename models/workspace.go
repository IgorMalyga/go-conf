package models

import "gorm.io/gorm"

type Workspace struct {
	gorm.Model
	CreatedByID uint
	CreatedBy   User `json:"createdBy" binding:"required" gorm:"foreignkey:CreatedByID"`
	Name        string
}
