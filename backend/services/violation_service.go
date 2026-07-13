package services

import (
	"github.com/MohamedFazil1406/RealTimeAlertSystem/models"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/repository"
)

type ViolationService struct {
	repo *repository.ViolationRepository
}

func NewViolationService() *ViolationService {

	return &ViolationService{
		repo: repository.NewViolationRepository(),
	}
}

func (s *ViolationService) GetHistory(

	vehicleID string,
	geofenceID string,
	startDate string,
	endDate string,
	limit int,

) ([]models.Violation, int64, error) {

	return s.repo.GetHistory(

		vehicleID,
		geofenceID,
		startDate,
		endDate,
		limit,
	)
}