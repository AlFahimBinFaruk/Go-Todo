package models

import (
	"time"

	"github.com/google/uuid"
)

type USER struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username  string
	Email     string `gorm:"uniqueIndex;not null"`
	Password  string `gorm:"not null"`
	Todos     []TODO `gorm:"foreignKey:UserId"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
