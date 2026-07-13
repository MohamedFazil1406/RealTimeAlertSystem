package repository

import (
	"github.com/MohamedFazil1406/RealTimeAlertSystem/database"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/models"
)

type LocationRepository struct{}

func NewLocationRepository() *LocationRepository {
	return &LocationRepository{}
}

func (r *LocationRepository) Save(location *models.Location) error {
	return database.DB.Create(location).Error
}

func (r *LocationRepository) GetLatest(vehicleID string) (*models.Location, error) {

	var location models.Location

	err := database.DB.
		Where("vehicle_id = ?", vehicleID).
		Order("timestamp desc").
		First(&location).Error

	return &location, err
}

func (r *LocationRepository) GetCurrentLocation(vehicleID string) (*models.Location, error) {

	var location models.Location

	err := database.DB.
		Where("vehicle_id = ?", vehicleID).
		Order("timestamp DESC").
		First(&location).Error

	return &location, err
}