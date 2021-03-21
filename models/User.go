package models

import (
	"time"

	_ "gorm.io/gorm"
)

type User struct {
	ID        int       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Name      string    `json:"name"`
	Pets      []*Pet    `json:"pets"`
}
