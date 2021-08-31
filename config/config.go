package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

func GetConfig() *Config {
	var config Config
	f, err := os.Open("config.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	return &config
}

type Config struct {
	Protocol   string `yaml:"protocol"`
	Hostname   string `yaml:"hostname"`
	Interval   int    `yaml:"interval"`
	ServerPort string `yaml:"server-port"`
}
