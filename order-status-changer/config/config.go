// This package the configuration for the order-adder service.
package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config is the configuration for the api server.
type Config struct {
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
		Port     string `yaml:"port"`
		SSLMode  string `yaml:"sslmode"`
	}
	Kafka struct {
		Brokers string `yaml:"brokers"`
	}
}

// InitConfig initializes the configuration for the api server.
func InitConfig() *Config {
	cfg := Config{}
	file, err := os.Open("config/cfg.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return &cfg
}
