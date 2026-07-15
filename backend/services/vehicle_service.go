package services

import (
	"github.com/MohamedFazil1406/RealTimeAlertSystem/dto"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/models"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/repository"
)

type VehicleService struct {
	repo *repository.VehicleRepository
}

func NewVehicleService() *VehicleService {
	return &VehicleService{
		repo: repository.NewVehicleRepository(),
	}
}

func (s *VehicleService) CreateVehicle(req dto.CreateVehicleRequest) (*models.Vehicle, error) {

	vehicle := models.Vehicle{
		VehicleNumber: req.VehicleNumber,
		DriverName:    req.DriverName,
		VehicleType:   req.VehicleType,
		Phone:         req.Phone,
	}

	err := s.repo.Create(&vehicle)

	return &vehicle, err
}

func (s *VehicleService) GetVehicles() ([]dto.VehicleWithLocation, error) {
	return s.repo.GetVehiclesWithLocation()
}