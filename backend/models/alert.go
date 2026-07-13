package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Alert struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	GeofenceID uuid.UUID

	VehicleID *uuid.UUID

	EventType string

	Status string `gorm:"default:active"`

	CreatedAt time.Time
}

func (a *Alert) BeforeCreate(tx *gorm.DB) error {
	a.ID = uuid.New()
	return nil
}