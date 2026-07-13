package utils

import (
	"encoding/json"

	"github.com/MohamedFazil1406/RealTimeAlertSystem/dto"
)

func ParsePolygon(data string) []Point {

	var coords []dto.Coordinate

	json.Unmarshal([]byte(data), &coords)

	points := []Point{}

	for _, c := range coords {

		points = append(points, Point{
			Lat: c.Latitude,
			Lng: c.Longitude,
		})

	}

	return points
}