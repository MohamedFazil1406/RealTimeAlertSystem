package services

import (
	"github.com/MohamedFazil1406/RealTimeAlertSystem/dto"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/models"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/repository"
	"github.com/google/uuid"
)

type AlertService struct {
	repo *repository.AlertRepository
}

func NewAlertService() *AlertService {
	return &AlertService{
		repo: repository.NewAlertRepository(),
	}
}

func (s *AlertService) Create(req dto.CreateAlertRequest) (*models.Alert, error) {
	println("AlertService.Create() called")

	geofenceID, err := uuid.Parse(req.GeofenceID)

	if err != nil {
		return nil, err
	}

	alert := models.Alert{
		GeofenceID: geofenceID,
		EventType:  req.EventType,
	}

	if req.VehicleID != "" {

		vehicleID, err := uuid.Parse(req.VehicleID)

		if err != nil {
			return nil, err
		}

		alert.VehicleID = &vehicleID
	}

	err = s.repo.Create(&alert)

	return &alert, err
}


func (s *AlertService) GetAll() ([]models.Alert, error) {
	return s.repo.GetAll()
}