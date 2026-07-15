package services

import (
	"log"
	"time"

	"github.com/MohamedFazil1406/RealTimeAlertSystem/dto"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/models"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/repository"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/utils"
	"github.com/MohamedFazil1406/RealTimeAlertSystem/websocket"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LocationService struct {
	locationRepo *repository.LocationRepository
	geoRepo      *repository.GeofenceRepository
	stateRepo *repository.VehicleStateRepository
	violationRepo *repository.ViolationRepository
	vehicleRepo *repository.VehicleRepository
}


func NewLocationService() *LocationService {

	return &LocationService{

	locationRepo: repository.NewLocationRepository(),

	geoRepo: repository.NewGeofenceRepository(),

	stateRepo: repository.NewVehicleStateRepository(),

	violationRepo: repository.NewViolationRepository(),
	vehicleRepo: repository.NewVehicleRepository(),
}
}

func (s *LocationService) UpdateLocation(req dto.UpdateLocationRequest) error {

	vehicleID, err := uuid.Parse(req.VehicleID)
	if err != nil {
		return err
	}

	timestamp, err := time.Parse(time.RFC3339, req.Timestamp)
	if err != nil {
		return err
	}

	location := models.Location{
		VehicleID: vehicleID,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Timestamp: timestamp,
	}

	// Save latest location
	if err := s.locationRepo.Save(&location); err != nil {
		return err
	}

	// Broadcast location update (for live map)
	websocket.Broadcast(gin.H{
		"type":       "location",
		"vehicle_id": vehicleID.String(),
		"latitude":   req.Latitude,
		"longitude":  req.Longitude,
		"timestamp":  req.Timestamp,
	})

	// Load all geofences
	geofences, err := s.geoRepo.GetAll()
	if err != nil {
		return err
	}

	currentPoint := utils.Point{
		Lat: req.Latitude,
		Lng: req.Longitude,
	}

	for _, geo := range geofences {

		polygon := utils.ParsePolygon(geo.Polygon)

		inside := utils.PointInPolygon(currentPoint, polygon)

		log.Printf("Geofence=%s Inside=%v", geo.Name, inside)

		state, err := s.stateRepo.Get(vehicleID, geo.ID)

		// First time this vehicle is checked against this geofence
		if err != nil {

			state = &models.VehicleState{
				VehicleID:  vehicleID,
				GeofenceID: geo.ID,
				IsInside:   inside,
			}

			if err := s.stateRepo.Save(state); err != nil {
				return err
			}

			continue
		}

		log.Printf("Previous=%v Current=%v", state.IsInside, inside)

		// No change
		if state.IsInside == inside {
			continue
		}

		event := "exit"
		if inside {
			event = "entry"
		}

		log.Printf("Vehicle %s detected", event)

		// Save violation
		violation := models.Violation{
			VehicleID:  vehicleID,
			GeofenceID: geo.ID,
			EventType:  event,
			Latitude:   req.Latitude,
			Longitude:  req.Longitude,
			Timestamp:  timestamp,
		}

		if err := s.violationRepo.Create(&violation); err != nil {
			return err
		}

		// Broadcast alert
		websocket.Broadcast(dto.AlertMessage{
			EventType:  event,
			VehicleID:  vehicleID.String(),
			GeofenceID: geo.ID.String(),
			Latitude:   req.Latitude,
			Longitude:  req.Longitude,
			Timestamp:  req.Timestamp,
		})

		// Update state
		state.IsInside = inside

		if err := s.stateRepo.Save(state); err != nil {
			return err
		}
	}

	return nil
}

func (s *LocationService) GetCurrentLocation(vehicleID string) (map[string]interface{}, error) {

	location, err := s.locationRepo.GetCurrentLocation(vehicleID)

	if err != nil {
		return nil, err
	}

	vehicle, err := s.vehicleRepo.GetByID(vehicleID)

	if err != nil {
		return nil, err
	}

	geofences, _ := s.geoRepo.GetAll()

	current := utils.Point{
		Lat: location.Latitude,
		Lng: location.Longitude,
	}

	var currentGeofences []gin.H

	for _, geo := range geofences {

		polygon := utils.ParsePolygon(geo.Polygon)

		if utils.PointInPolygon(current, polygon) {

			currentGeofences = append(currentGeofences, gin.H{
				"geofence_id": geo.ID,
				"geofence_name": geo.Name,
				"category": geo.Category,
			})
		}
	}

	response := map[string]interface{}{
		"vehicle_id": vehicle.ID,
		"vehicle_number": vehicle.VehicleNumber,
		"current_location": gin.H{
			"latitude": location.Latitude,
			"longitude": location.Longitude,
			"timestamp": location.Timestamp,
		},
		"current_geofences": currentGeofences,
	}

	return response, nil
}