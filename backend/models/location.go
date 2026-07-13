package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Location struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	VehicleID uuid.UUID

	Latitude  float64
	Longitude float64

	Timestamp time.Time

	CreatedAt time.Time
}

func (l *Location) BeforeCreate(tx *gorm.DB) error {
	l.ID = uuid.New()
	return nil
}