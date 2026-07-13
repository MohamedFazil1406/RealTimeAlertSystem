package repository

import (
	"github.com/MohamedFazil1406/RealTimeAlertSystem/database"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/models"
	"github.com/google/uuid"
)

type VehicleStateRepository struct{}

func NewVehicleStateRepository() *VehicleStateRepository {
	return &VehicleStateRepository{}
}

func (r *VehicleStateRepository) Get(vehicleID, geofenceID uuid.UUID) (*models.VehicleState, error) {

	var state models.VehicleState

	err := database.DB.
		Where("vehicle_id=? AND geofence_id=?", vehicleID, geofenceID).
		First(&state).Error

	return &state, err
}

func (r *VehicleStateRepository) Save(state *models.VehicleState) error {

	return database.DB.Save(state).Error
}