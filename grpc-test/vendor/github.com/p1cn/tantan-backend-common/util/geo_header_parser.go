package util

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrSyntax      = fmt.Errorf("Geolocation: Syntax Error")
	ErrBadLocation = fmt.Errorf("Geolocation: Bad Location")
)

type GeoUri struct {
	Latitude  float64
	Longitude float64
	CordC     *float64 `json:"CordC,omitempty"`

	ss     string
	parsed string
}

// IsDefaultLocation returns default if this is a "default" location.
// "default" locations are almost certainly reported because a proper
// location could not be fetched, and not because the user actually was
// located there.
//
// Currently three locations are listed, (0,0) (1,1) and (30, 104).
// (30, 104) is the center of china.
func (self *GeoUri) IsDefaultLocation() bool {
	return self.Longitude == 0 && self.Latitude == 0 ||
		self.Longitude == 1 && self.Latitude == 1 ||
		self.Longitude == 104 && self.Latitude == 30
}

// IsBadLocation returns true if this is a known bad location.
// Currently one bad location is listed: (-0.0027, -0.01)
func (self *GeoUri) IsBadLocation() bool {
	return self.Longitude == -0.01 && self.Latitude == -0.0027
}

// IsInsideChina returns true if this is inside China region
// TODO: improve accuracy in China boundary detection
func (self *GeoUri) IsInsideChina() bool {
	if self.Latitude < 0.8293 || self.Latitude > 55.8271 {
		return false
	}
	if self.Longitude < 72.004 || self.Longitude > 137.8347 {
		return false
	}
	return true
}

func (self *GeoUri) Parse(s string) error {
	self.ss = s
	return self.parseGeoUri()
}

func (self *GeoUri) parseGeoUri() error {
	err := self.parseGeoScheme()
	if err != nil {
		return err
	}
	err = self.parseGeoPath()
	if err != nil {
		return err
	}
	return nil
}

func (self *GeoUri) parseGeoScheme() error {
	ok := self.toKey("geo:")
	if !ok {
		return ErrSyntax
	}
	if len(self.parsed) != 0 {
		return ErrSyntax
	}
	return nil
}

func (self *GeoUri) parseGeoPath() error {
	var err error
	self.Latitude, err = self.number()
	if err != nil {
		return ErrSyntax
	}
	if len(self.ss) > 0 {
		if self.ss[0] != ',' {
			return ErrSyntax
		}
	} else {
		return ErrSyntax
	}
	self.ss = self.ss[1:]
	self.Longitude, err = self.number()
	if err != nil {
		return ErrSyntax
	}
	if self.Longitude > 180 || self.Longitude < -180 || self.Latitude > 90 || self.Latitude < -90 {
		return ErrBadLocation
	}

	return nil
}

func (self *GeoUri) toKey(key string) bool {
	split := strings.SplitN(self.ss, key, 2)
	if len(split) == 0 {
		self.ss = ""
		self.parsed = ""
		return false
	}
	if split[0] == self.ss {
		return false
	}
	if len(split) == 1 {
		self.ss = split[0]
	} else {
		self.ss = split[1]
		self.parsed = split[0]
	}
	return true
}

func (self *GeoUri) number() (float64, error) {
	foundPoint := false
	var i int
	var c rune
	for i, c = range self.ss {
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		case '.':
			if foundPoint {
				f, err := strconv.ParseFloat(self.ss[0:i], 64)
				self.ss = self.ss[i:]
				return f, err
			}
			foundPoint = true
		case '-':
			if i != 0 {
				return 0, ErrSyntax
			}
		default:
			if i == 0 {
				return 0, ErrSyntax
			}
			f, err := strconv.ParseFloat(self.ss[0:i], 64)
			self.ss = self.ss[i:]
			return f, err
		}
	}
	f, err := strconv.ParseFloat(self.ss[0:], 64)
	self.ss = self.ss[i:]
	return f, err
}
