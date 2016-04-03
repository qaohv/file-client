package main

import (
	"encoding/json"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
)

type Config struct {
	Host              string `json:"host"`
	Port              string `json:"port"`
	DefaultPathToSave string `json:"default-path-to-save"`
}

func ReadConfig(path string) *Config {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.WithField("error", err.Error()).Error("Can't read file by path " + path)
		return nil
	}

	config := new(Config)

	if err := json.Unmarshal(file, config); err != nil {
		log.WithField("error", err.Error()).Error("Can't unmarshal file " + path)
		return nil
	}

	return config
}
