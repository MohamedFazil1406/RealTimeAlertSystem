package utils

import "github.com/MohamedFazil1406/RealTimeAlertSystem/dto"

func ValidCategory(category string) bool {
	valid := map[string]bool{
		"delivery_zone":   true,
		"restricted_zone": true,
		"toll_zone":       true,
		"customer_area":   true,
	}

	return valid[category]
}

func ValidCoordinates(coords []dto.Coordinate) bool {
	if len(coords) < 4 {
		return false
	}

	for _, c := range coords {
		if c.Latitude < -90 || c.Latitude > 90 {
			return false
		}

		if c.Longitude < -180 || c.Longitude > 180 {
			return false
		}
	}

	return true
}