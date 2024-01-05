package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Token string `json:"token"`
}

func LoadConfig(filename string) Config {
	var config Config

	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("Unable to parse file: %v", err)
	}

	return config
}
