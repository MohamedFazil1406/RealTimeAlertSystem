package repository

import (
	"github.com/MohamedFazil1406/RealTimeAlertSystem/database"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/models"
)

type GeofenceRepository struct{}

func NewGeofenceRepository() *GeofenceRepository {
	return &GeofenceRepository{}
}

func (r *GeofenceRepository) Create(geo *models.Geofence) error {
	return database.DB.Create(geo).Error
}

func (r *GeofenceRepository) GetAll() ([]models.Geofence, error) {

	var geofences []models.Geofence

	err := database.DB.Find(&geofences).Error

	return geofences, err
}