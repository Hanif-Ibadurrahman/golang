package models

import "gorm.io/gorm"

type Member struct {
	gorm.Model
	Name   string `gorm:"type:varchar(255)" json:"name"`
	Email  string `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	UserID uint64 `gorm:"not null" json:"-"`
	User   User   `gorm:"foreignkey:UserID;references:user;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
}
