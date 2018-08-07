package util

import (
	"math"
)

const (
	DegToRad    = 0.01745329251
	EarthRadius = 6371
)

//Don't know if this is correct
func CalculateDistanceInKm(lo1 float64, la1 float64, lo2 float64, la2 float64) float64 {
	dlat := (la2 - la1) * DegToRad
	dlon := (lo2 - lo1) * DegToRad
	lat1 := la1 * DegToRad
	lat2 := la2 * DegToRad
	a := math.Sin(dlat/2)*math.Sin(dlat/2) + math.Sin(dlon/2)*math.Sin(dlon/2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return EarthRadius * c // should be rounded to 4 digit precision
}
