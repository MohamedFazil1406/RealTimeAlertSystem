package utils

func PointInPolygon(point Point, polygon []Point) bool {

	inside := false

	j := len(polygon) - 1

	for i := 0; i < len(polygon); i++ {

		intersect :=
			((polygon[i].Lng > point.Lng) != (polygon[j].Lng > point.Lng)) &&
				(point.Lat <
					(polygon[j].Lat-polygon[i].Lat)*
						(point.Lng-polygon[i].Lng)/
						(polygon[j].Lng-polygon[i].Lng)+
						polygon[i].Lat)

		if intersect {
			inside = !inside
		}

		j = i
	}

	return inside
}