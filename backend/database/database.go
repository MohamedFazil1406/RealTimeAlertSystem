package database

import (
	"fmt"
	"log"

	"github.com/MohamedFazil1406/RealTimeAlertSystem/config"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectDatabase() {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.GetEnv("DB_HOST"),
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASSWORD"),
		config.GetEnv("DB_NAME"),
		config.GetEnv("DB_PORT"),
		config.GetEnv("DB_SSLMODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	DB = db

	err = DB.AutoMigrate(
	&models.Vehicle{},
	&models.Geofence{},
	&models.Location{},
	&models.Alert{},
	&models.Violation{},
	&models.VehicleState{},
)

if err != nil {
	log.Fatal(err)
}

	log.Println("Database Connected Successfully")
}