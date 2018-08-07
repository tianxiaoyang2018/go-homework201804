package config

import (
	"encoding/json"
	"io/ioutil"
)

func UnmarshalJsonConfig(file string, obj interface{}) error {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, obj)
	if err != nil {
		return err
	}

	// TODO: Check if valid

	return nil
}
