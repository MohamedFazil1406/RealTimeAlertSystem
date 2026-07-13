package dto

type AlertMessage struct {

	EventType string `json:"event_type"`

	VehicleID string `json:"vehicle_id"`

	GeofenceID string `json:"geofence_id"`

	Latitude float64 `json:"latitude"`

	Longitude float64 `json:"longitude"`

	Timestamp string `json:"timestamp"`
}