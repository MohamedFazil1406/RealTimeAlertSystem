package dto

type Coordinate struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type CreateGeofenceRequest struct {
	Name        string       `json:"name" binding:"required"`
	Description string       `json:"description"`
	Coordinates []Coordinate `json:"coordinates" binding:"required"`
	Category    string       `json:"category" binding:"required"`
}