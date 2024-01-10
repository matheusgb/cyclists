package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Database Database
	Api      Api
}

type Database struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

type Api struct {
	Port string
}

func Init() Config {
	var configJson Config
	configFile, err := os.Open("./config/config.json")
	if err != nil {
		log.Fatal("Config file not found: ", err)
	}

	defer configFile.Close()

	jsonDecoder := json.NewDecoder(configFile)
	err = jsonDecoder.Decode(&configJson)
	if err != nil {
		log.Fatal("Error parsing config file: ", err)
	}

	return configJson
}
