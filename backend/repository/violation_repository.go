package repository

import (
	"time"

	"github.com/MohamedFazil1406/RealTimeAlertSystem/database"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/models"
)

type ViolationRepository struct{}

func NewViolationRepository() *ViolationRepository {
	return &ViolationRepository{}
}

func (r *ViolationRepository) Create(v *models.Violation) error {
	return database.DB.Create(v).Error
}

func (r *ViolationRepository) GetHistory(
	vehicleID string,
	geofenceID string,
	startDate string,
	endDate string,
	limit int,
) ([]models.Violation, int64, error) {

	var violations []models.Violation

	query := database.DB.Model(&models.Violation{})

	if vehicleID != "" {
		query = query.Where("vehicle_id = ?", vehicleID)
	}

	if geofenceID != "" {
		query = query.Where("geofence_id = ?", geofenceID)
	}

	if startDate != "" {

		t, _ := time.Parse(time.RFC3339, startDate)

		query = query.Where("timestamp >= ?", t)
	}

	if endDate != "" {

		t, _ := time.Parse(time.RFC3339, endDate)

		query = query.Where("timestamp <= ?", t)
	}

	var total int64

	query.Count(&total)

	err := query.
		Order("timestamp DESC").
		Limit(limit).
		Find(&violations).Error

	return violations, total, err
}