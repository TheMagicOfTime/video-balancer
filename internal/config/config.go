package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port    string
	CDNHost string
}

func LoadConfig() (*Config, error) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "50051"
	}

	cdnHost := os.Getenv("CDN_HOST")
	if cdnHost == "" {
		return nil, fmt.Errorf("CDN_HOST environment variable is not set")
	}

	return &Config{
		Port:    port,
		CDNHost: cdnHost,
	}, nil
}
