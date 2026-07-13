package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Vehicle struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	VehicleNumber string    `gorm:"unique;not null"`
	DriverName    string    `gorm:"not null"`
	VehicleType   string    `gorm:"not null"`
	Phone         string    `gorm:"not null"`
	Status        string    `gorm:"default:active"`

	CreatedAt time.Time
	UpdatedAt time.Time
}



func (v *Vehicle) BeforeCreate(tx *gorm.DB) error {
	v.ID = uuid.New()
	return nil
}