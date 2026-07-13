package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Violation struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	VehicleID uuid.UUID

	GeofenceID uuid.UUID

	EventType string

	Latitude  float64
	Longitude float64

	Timestamp time.Time

	CreatedAt time.Time
}

func (v *Violation) BeforeCreate(tx *gorm.DB) error {
	v.ID = uuid.New()
	return nil
}