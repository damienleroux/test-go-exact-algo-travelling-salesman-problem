package mathUtils

import (
	"math"
)

func Factorial(x int) int {
	if x == 0 {
		return 1
	}

	return x * Factorial(x-1)
}

// returns the mathematical distance between two geo points (lat + long)
// the returned value is either "m" for meter or "" for a raw result.
// this func aims to compare several distance from different points
func Distance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit string) float64 {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}
	dist = math.Acos(dist)
	if unit == "m" {
		dist = dist * 180 / PI
		dist = dist * 60 * 1.1515

		dist = dist * 1.609344 * 1000 //m
	}

	return dist
}

// returns the mathematical distance between two geo points (lat + long)
// this func aims to compare several distance from different points
func RawDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64) float64 {
	return Distance(lat1, lng1, lat2, lng2, "")
}
