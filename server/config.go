package main

import (
	"os"
)

type Config struct {
	HostPort	string
}

func InitConfig() *Config {
	// Connection config
	HostPort, exists := os.LookupEnv("HOST_PORT")
	if exists == false {
		HostPort = "8080"
	}
	return &Config{
		HostPort: HostPort,
	}
}
