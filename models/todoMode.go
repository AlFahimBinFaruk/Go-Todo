package models

import (
	"time"
)

type TODO struct {
	ID        uint `gorm:primaryKey`
	Title     string
	Desc      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
