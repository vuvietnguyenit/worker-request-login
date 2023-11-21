package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func login(username string, password string) {
	// Create body data
	bodyData := map[string]string{"username": username, "password": password}
	bodyReader, err := json.Marshal(bodyData)
	if err != nil {
		log.Printf("error when parse request data: %s\n", err)
		return
	}
	req, err := http.NewRequest(http.MethodPost, configData.Api.Login, bytes.NewBuffer(bodyReader))
	if err != nil {
		log.Printf("client: could not create request: %s\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("client: error making http request: %s\n", err)
		return
	}
	log.Printf("login with username: %s, password: %s, status_code: %d\n", username, password, res.StatusCode)

}

func main() {
	log.Println("worker started.")
	log.Println("start read config file...")
	readConfigFile("config.yaml")
	data := readDataFile("users_with_password_dummy.csv")
	for {
		row := pickRandomElementIn2dArray(data)
		interval := pickRandomInterval()
		username := row[2]
		password := row[5]
		login(username, password)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}
