package main

import (
	"os"
)

type Config struct {
	HostPort	string
}

var config = Config{}

func InitConfig() {
	config = Config{}

	// Connection config
	HostPort, exists := os.LookupEnv("HOST_PORT")
	if exists == false {
		HostPort = "8080"
	}
	config.HostPort = HostPort
}
