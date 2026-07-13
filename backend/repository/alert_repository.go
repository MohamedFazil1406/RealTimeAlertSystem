package repository

import (
	"github.com/MohamedFazil1406/RealTimeAlertSystem/database"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/models"
)

type AlertRepository struct{}

func NewAlertRepository() *AlertRepository {
	return &AlertRepository{}
}

func (r *AlertRepository) Create(alert *models.Alert) error {
	return database.DB.Create(alert).Error
}

func (r *AlertRepository) GetAll() ([]models.Alert, error) {

	var alerts []models.Alert

	err := database.DB.Find(&alerts).Error

	return alerts, err
}