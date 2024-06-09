package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/matheusgb/cyclists/src/config"
	requests "github.com/matheusgb/cyclists/src/controllers/requests/user"
	"github.com/matheusgb/cyclists/src/layers"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Current dir: %v\n", currentDir)

	configInit := config.Init()
	configFile, err := os.Open("../../../config/config.json")
	if err != nil {
		log.Fatal("config file not found: ", err)
	}

	defer configFile.Close()

	jsonDecoder := json.NewDecoder(configFile)
	err = jsonDecoder.Decode(&configInit)
	if err != nil {
		log.Fatal("error parsing config file: ", err)
	}

	config.InitializedConfigs = configInit
	config.InitializedConfigs.Enviroment = "test"

	code := m.Run()
	os.Exit(code)
}

func TestCreateUser(t *testing.T) {
	request := requests.CreateUser{
		Name:                 "test",
		Email:                "iuasdhfiaf@hsdiaufh.iqu",
		Password:             "123456",
		PasswordConfirmation: "123456",
	}

	marshalled, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("impossible to marshall request: %s", err)
	}

	app := layers.Setup(config.InitializedConfigs)
	req, _ := http.NewRequest(
		"POST",
		"/api/v1/register",
		bytes.NewReader(marshalled),
	)
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("Error on test: %v", err)
	}

	var resBodyMap map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&resBodyMap)
	if err != nil {
		t.Fatalf("Error on test: %v", err)
	}

	assert.EqualValues(t, 201, res.StatusCode)

	req, _ = http.NewRequest(
		"DELETE",
		fmt.Sprintf("/api/v1/user/%s", resBodyMap["id"].(string)),
		nil,
	)

	fmt.Println(resBodyMap["id"].(string))

	res, err = app.Test(req, -1)
	if err != nil {
		t.Fatalf("Error on test: %v", err)
	}

	assert.EqualValues(t, 204, res.StatusCode)
}
