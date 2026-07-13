package services

import (
	"encoding/json"

	"github.com/MohamedFazil1406/RealTimeAlertSystem/dto"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/models"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/repository"
)

type GeofenceService struct {
	repo *repository.GeofenceRepository
}

func NewGeofenceService() *GeofenceService {
	return &GeofenceService{
		repo: repository.NewGeofenceRepository(),
	}
}

func (s *GeofenceService) Create(req dto.CreateGeofenceRequest) (*models.Geofence, error) {

	polygon, _ := json.Marshal(req.Coordinates)

	geo := models.Geofence{
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
		Polygon:     string(polygon),
	}

	err := s.repo.Create(&geo)

	return &geo, err
}

func (s *GeofenceService) GetAll() ([]models.Geofence, error) {
	return s.repo.GetAll()
}