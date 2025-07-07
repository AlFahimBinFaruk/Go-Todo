package models

import (
	"time"

	"github.com/google/uuid"
)

type TODO struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title     string
	Desc      string
	UserId    uuid.UUID `gorm:"type:uuid;index"` // foreign key
	CreatedAt time.Time
	UpdatedAt time.Time
}
