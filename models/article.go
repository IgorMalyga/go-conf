package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title        string
	CreatedBy    User
	UsersCanEdit []User
	Tags         []Tag
	Workspace    Workspace
}

type Tag struct {
	gorm.Model
	Name string
}
