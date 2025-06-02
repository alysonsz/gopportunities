package models

import (
	"time"

	"gorm.io/gorm"
)

// Opportunity representa uma vaga ou oportunidade cadastrada no sistema
type Opportunity struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Title       string         `json:"title" validate:"required,min=3,max=100"`
	Description string         `json:"description" validate:"required"`
	Location    string         `json:"location"`
	Status      string         `json:"status" validate:"required,oneof=open closed"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
