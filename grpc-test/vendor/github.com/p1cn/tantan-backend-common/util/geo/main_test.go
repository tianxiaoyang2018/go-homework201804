package geo

import (
	"flag"
	"log"
	"os"
	"testing"
)

var configPath = flag.String("config", "/etc/putong/config.json", "Path to configuration file.")

var polygons Polygons

func TestMain(m *testing.M) {
	var err error
	polygons, err = FromJSONFile("./test/usa-and-korea.geo.json")
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}
