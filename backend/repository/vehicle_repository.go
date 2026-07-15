package repository

import (
	"log"

	"github.com/MohamedFazil1406/RealTimeAlertSystem/database"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/dto"
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

func (r *VehicleRepository) GetVehiclesWithLocation() ([]dto.VehicleWithLocation, error) {

	var vehicles []models.Vehicle

	if err := database.DB.Find(&vehicles).Error; err != nil {
		return nil, err
	}

	var result []dto.VehicleWithLocation

	for _, v := range vehicles {

		var location models.Location

		err := database.DB.
			Where("vehicle_id = ?", v.ID).
			Order("timestamp DESC").
			First(&location).Error

					if err != nil {
			log.Printf("No location found for vehicle %s : %v", v.VehicleNumber, err)
		} else {
			log.Printf(
				"Vehicle %s -> Lat=%f Lng=%f",
				v.VehicleNumber,
				location.Latitude,
				location.Longitude,
			)
		}

		item := dto.VehicleWithLocation{
			ID:            v.ID,
			VehicleNumber: v.VehicleNumber,
			DriverName:    v.DriverName,
			VehicleType:   v.VehicleType,
			Phone:         v.Phone,
			Status:        v.Status,
		}

		if err == nil {
			item.Latitude = &location.Latitude
			item.Longitude = &location.Longitude
		}

		result = append(result, item)
	}

	return result, nil
}