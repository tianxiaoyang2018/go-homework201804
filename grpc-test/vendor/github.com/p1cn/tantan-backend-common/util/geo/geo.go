package geo

import (
	"encoding/json"
	"os"

	"github.com/p1cn/tantan-backend-common/util"

	ggeo "github.com/kellydunn/golang-geo"
)

type Polygons []*ggeo.Polygon

func (ps *Polygons) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	p, err := FromJSONFile(s)
	if err != nil {
		return err
	}
	*ps = p
	return nil
}

func (ps Polygons) Contains(coor *util.GeoUri) bool {
	if coor == nil {
		return true
	}
	point := ggeo.NewPoint(coor.Latitude, coor.Longitude)
	for _, p := range ps {
		if p.Contains(point) {
			return true
		}
	}
	return false
}

type geojson struct {
	Features []struct {
		Geometry struct {
			Type        string          `json:"type"`
			Coordinates json.RawMessage `json:"coordinates"`
		} `json:"geometry"`
	} `json:"features"`
}

type PolygonCoor [][][2]float64
type MultiPolygonCoor [][][][2]float64

func FromJSONFile(filepath string) (Polygons, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	collection := new(geojson)
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&collection); err != nil {
		return nil, err
	}
	ps := make([]*ggeo.Polygon, 0)
	for _, feature := range collection.Features {
		switch feature.Geometry.Type {
		case "Polygon":
			var coor PolygonCoor
			if err := json.Unmarshal(feature.Geometry.Coordinates, &coor); err != nil {
				return nil, err
			}
			p := &ggeo.Polygon{}
			if len(coor) != 1 {
				continue
			}
			for _, point := range coor[0] {
				p.Add(ggeo.NewPoint(point[1], point[0]))

			}
			ps = append(ps, p)
		case "MultiPolygon":
			var coors MultiPolygonCoor
			if err := json.Unmarshal(feature.Geometry.Coordinates, &coors); err != nil {
				return nil, err
			}
			for _, coor := range coors {
				p := &ggeo.Polygon{}
				if len(coor) != 1 {
					continue
				}
				for _, point := range coor[0] {
					p.Add(ggeo.NewPoint(point[1], point[0]))

				}
				ps = append(ps, p)
			}
		}
	}
	return ps, nil
}
