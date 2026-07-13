package dto

type CreateAlertRequest struct {
	GeofenceID string `json:"geofence_id" binding:"required"`
	VehicleID  string `json:"vehicle_id"`
	EventType  string `json:"event_type" binding:"required"`
}