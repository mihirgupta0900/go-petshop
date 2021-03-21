package models

import (
	"time"

	_ "gorm.io/gorm"
)

type Pet struct {
	ID        int       `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Name      string    `json:"name"`
	Age       uint16    `json:"age"` // u signifies non-negative
	Breed     string    `json:"breed"`
	UserID    int       `json:"user_id"`
}
