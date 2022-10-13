package models

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
