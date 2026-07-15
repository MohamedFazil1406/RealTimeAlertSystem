package services

import (
	"github.com/MohamedFazil1406/RealTimeAlertSystem/dto"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/repository"
)

type DashboardService struct {
	repo *repository.DashboardRepository
}

func NewDashboardService() *DashboardService {
	return &DashboardService{
		repo: repository.NewDashboardRepository(),
	}
}

func (s *DashboardService) GetStats() (*dto.DashboardStats, error) {
	return s.repo.GetStats()
}