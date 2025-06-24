package models

import (
	"time"

	"github.com/google/uuid"
)

type TODO struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title     string
	Desc      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
