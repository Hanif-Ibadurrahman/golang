package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqetype:varchar(255)" json:"username"`
	Password string `gorm:"->;<-;not null" json:"-"`
	IsStaff  bool   `gorm:"not null" json:"-"`
	Token    string `gorm:"-" json:"token,omitempty"`
}
