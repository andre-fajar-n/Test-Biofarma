package utils

import (
	"math"
)

// DistanceByLatLon calculate distance between 2 places based on latitude and longitude.
// This calculation use Haversine Method
func DistanceByLatLon(lat1, lon1, lat2, lon2 float64) float64 {
	R := float64(6371000)

	phi1 := lat1 * math.Pi / 180

	phi2 := lat2 * math.Pi / 180

	deltaPhi := (lat2 - lat1) * math.Pi / 180

	deltaLambda := (lon2 - lon1) * math.Pi / 180

	a := math.Pow(math.Sin(deltaPhi/2), 2) + math.Cos(phi1)*math.Cos(phi2)*math.Pow(math.Sin(deltaLambda/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}
