package geo

import (
	"testing"

	"github.com/p1cn/tantan-backend-common/util"
)

// https://geojson-maps.ash.ms
func TestParsePolygon(t *testing.T) {
	point := util.GeoUri{Latitude: 37.0902, Longitude: -95.7129}
	if !polygons.Contains(&point) {
		t.Log("not in")
	} else {
		t.Log("in")
	}
}

// @TODO
func BenchmarkFContains(b *testing.B) {
	for i := 0; i < b.N; i++ {
		point := util.GeoUri{Latitude: -37.0902, Longitude: -95.7129}
		polygons.Contains(&point)
	}
}
