package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"unique;not null"`
	Password string `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(bytes)
	return
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
