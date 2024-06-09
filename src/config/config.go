package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Database   Database
	Api        Api
	Jwt        Jwt
	SendGrid   SendGrid
	Enviroment string
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

type Jwt struct {
	Secret string
}

type SendGrid struct {
	ApiKey      string
	FromEmail   string
	URLFrontEnd string
}

func Init() *Config {
	return &Config{}
}

var InitializedConfigs *Config

func (config *Config) MountConfigs() {
	configFile, err := os.Open("./src/config/config.json")
	if err != nil {
		log.Fatal("config file not found: ", err)
	}

	defer configFile.Close()

	jsonDecoder := json.NewDecoder(configFile)
	err = jsonDecoder.Decode(&config)
	if err != nil {
		log.Fatal("error parsing config file: ", err)
	}

	InitializedConfigs = config
}
