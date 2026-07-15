package repository

import (
	"github.com/MohamedFazil1406/RealTimeAlertSystem/database"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/dto"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/models"
)

type DashboardRepository struct{}

func NewDashboardRepository() *DashboardRepository {
	return &DashboardRepository{}
}

func (r *DashboardRepository) GetStats() (*dto.DashboardStats, error) {

	stats := &dto.DashboardStats{}

	database.DB.Model(&models.Vehicle{}).Count(&stats.TotalVehicles)

	database.DB.Model(&models.Geofence{}).Count(&stats.TotalGeofences)

	database.DB.Model(&models.Violation{}).Count(&stats.TotalViolations)

	// For now alerts = violations
	stats.TotalAlerts = stats.TotalViolations

	return stats, nil
}