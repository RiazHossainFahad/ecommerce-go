package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version  string
	AppName  string
	HttpPort int
}

var configurations Config

func loadConfig() {
	err := godotenv.Load()

	if err != nil {
		// Log a warning and continue, or log.Fatal if the file is required
		log.Fatal("Warning: Error loading .env file\nError: ", err)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		log.Fatal("Env error: VERSION is missing")
	}

	appName := os.Getenv("APP_NAME")
	if appName == "" {
		log.Fatal("Env error: APP_NAME is missing")
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Fatal("Env error: HTTP_PORT is missing")
	}

	actualPort, err := strconv.Atoi(httpPort)
	if err != nil {
		log.Fatal("Env error: Invalid HTTP_PORT given\nError: ", err)
	}

	configurations = Config{
		Version:  version,
		AppName:  appName,
		HttpPort: actualPort,
	}
}

func GetConfig() Config {
	loadConfig()

	return configurations
}
