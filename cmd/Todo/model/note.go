package model

import (
	"gorm.io/gorm"
)

// Product struct
type Note struct {
	gorm.Model
	Content     string `gorm:"not null" json:"content"`
	IsCompleted bool   `json:"isCompleted"`
}
