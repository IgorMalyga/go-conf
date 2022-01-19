package models

import "gorm.io/gorm"

type Workspace struct {
	gorm.Model
	CreatedBy User
	Members   []User
}
