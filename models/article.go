package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title        string
	CreatedByID  uint
	CreatedBy    User
	UsersCanEdit []*User `gorm:"many2many:user_articles;"`
	Tags         []Tag   `gorm:"many2many:tag_articles;"`
	WorkspaceID  uint
	Workspace    Workspace
}

type Tag struct {
	gorm.Model
	Name string
}
