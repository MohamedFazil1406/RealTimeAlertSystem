package dto

type DashboardStats struct {
	TotalVehicles   int64 `json:"total_vehicles"`
	TotalGeofences  int64 `json:"total_geofences"`
	TotalViolations int64 `json:"total_violations"`
	TotalAlerts     int64 `json:"total_alerts"`
}