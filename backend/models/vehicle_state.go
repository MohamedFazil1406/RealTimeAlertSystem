package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type VehicleState struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`

	VehicleID uuid.UUID

	GeofenceID uuid.UUID

	IsInside bool

	UpdatedAt time.Time
}

func (v *VehicleState) BeforeCreate(tx *gorm.DB) error {

	v.ID = uuid.New()

	return nil
}