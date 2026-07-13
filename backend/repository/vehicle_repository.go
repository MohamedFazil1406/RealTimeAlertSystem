package repository

import (
	"github.com/MohamedFazil1406/RealTimeAlertSystem/database"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/models"
)

type VehicleRepository struct{}

func NewVehicleRepository() *VehicleRepository {
	return &VehicleRepository{}
}

func (r *VehicleRepository) Create(vehicle *models.Vehicle) error {
	return database.DB.Create(vehicle).Error
}

func (r *VehicleRepository) GetAll() ([]models.Vehicle, error) {

	var vehicles []models.Vehicle

	err := database.DB.Find(&vehicles).Error

	return vehicles, err
}

func (r *VehicleRepository) GetByID(id string) (*models.Vehicle, error) {

	var vehicle models.Vehicle

	err := database.DB.
		Where("id = ?", id).
		First(&vehicle).Error

	return &vehicle, err
}