package util

import "testing"

func TestValidCoordinates(t *testing.T) {
	validCoords := [][]float64{
		{40.01536, -83.083089},
		{89.924320, -179.342091},
		{90, -180},
	}
	for _, coord := range validCoords {
		if !ValidCoordinates(coord) {
			t.Errorf("%+v is a valid coordinate", coord)
		}
	}
	invalidCoords := [][]float64{
		{40.01536, -203.980261},
		{101.980324, -179.342091},
		{-123.980261, 180.098734},
	}
	for _, coord := range invalidCoords {
		if ValidCoordinates(coord) {
			t.Errorf("%+v is an invalid coordinate", coord)
		}
	}
}

func TestValidLatitude(t *testing.T) {
	validLats := []float64{40.01536, 0.805079, -83.083089, 89.924320, -89.342091, 70, 90, -90}
	for _, lat := range validLats {
		if !ValidLatitude(lat) {
			t.Errorf("%f is a valid latitude", lat)
		}
	}
	invalidLats := []float64{91, 90.098734, 101.980324, -91, -123.980261}
	for _, lat := range invalidLats {
		if ValidLatitude(lat) {
			t.Errorf("%f is not a valid latitude", lat)
		}
	}
}

func TestValidLongitude(t *testing.T) {
	validLons := []float64{40.01536, 0.805079, -83.083089, 179.924320, -179.342091, 160, 180, -180}
	for _, lon := range validLons {
		if !ValidLongitude(lon) {
			t.Errorf("%f is a valid longitude", lon)
		}
	}
	invalidLons := []float64{181, 180.098734, 191.980324, -181, -203.980261}
	for _, lon := range invalidLons {
		if ValidLongitude(lon) {
			t.Errorf("%f is an invalid longitude", lon)
		}
	}
}
