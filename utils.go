package main

import (
	"encoding/csv"
	"gopkg.in/yaml.v3"
	"log"
	"math/rand"
	"os"
)

type Config struct {
	Api struct {
		Login string `yaml:"login"`
	} `yaml:"api"`
	Config struct {
		Interval []float32 `yaml:"interval"`
	} `yaml:"config"`
}

func pickRandomInterval() float32 {
	randomIndex := rand.Intn(len(configData.Config.Interval))
	pick := configData.Config.Interval[randomIndex]
	return pick
}

func pickRandomElementIn2dArray(data [][]string) []string {
	randomIndex := rand.Intn(len(data))
	pick := data[randomIndex]
	return pick

}

func readDataFile(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	return data
}

func readConfigFile(filename string) {
	config := Config{}
	f, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error when read config file: %v", err.Error())
	}
	err = yaml.Unmarshal(f, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	configData = &config
}
