package config

import (
	"os"
)

type Config struct {
	HostPort	string
}

var config = Config{}

func Init() {
	config = Config{}

	// Connection config
	HostPort, exists := os.LookupEnv("HOST_PORT")
	if exists == false {
		HostPort = "8080"
	}
	config.HostPort = HostPort
}

func Get() Config {
	if &config == nil {
		Init()
	}
	return config
}
