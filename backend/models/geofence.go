package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Geofence struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name        string    `gorm:"not null"`
	Description string
	Category    string `gorm:"not null"`

	Polygon string `gorm:"type:text"`

	Status string `gorm:"default:active"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (g *Geofence) BeforeCreate(tx *gorm.DB) error {
	g.ID = uuid.New()
	return nil
}