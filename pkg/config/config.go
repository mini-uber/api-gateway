package config

import (
	"os"
)

type Config struct {
	GATEWAY_PORT   string
	AUTH_ADDR      string
	PASSENGER_ADDR string
	DRIVER_ADDR    string
}

var LoadedConfig *Config

func LoadConfig() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	authAddr := os.Getenv("AUTH_ADDR")
	passengerAddr := os.Getenv("PASSENGER_ADDR")
	driverAddr := os.Getenv("DRIVER_ADDR")
	
	LoadedConfig = &Config{
		GATEWAY_PORT:   port,
		AUTH_ADDR:      authAddr,
		PASSENGER_ADDR: passengerAddr,
		DRIVER_ADDR:    driverAddr,
	}
}
