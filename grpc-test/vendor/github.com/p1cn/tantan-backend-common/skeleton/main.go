package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/p1cn/tantan-backend-common/skeleton/config"
	"github.com/p1cn/tantan-backend-common/skeleton/template"
)

var configFile = flag.String("config", "", "config file path")

func main() {

	flag.Parse()
	log.SetFlags(log.Lshortfile)
	data, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}

	cfg := new(config.Config)
	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		log.Fatal(err)
	}

	tplVar := template.ModelFromConfig(cfg)

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = template.GenerateTemplates(dir, *tplVar)
	if err != nil {
		log.Fatal(err)
	}
}
