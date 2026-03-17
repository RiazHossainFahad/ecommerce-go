package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version   string
	AppName   string
	HttpPort  int
	JwtSecret string

	DbHost     string
	DbPort     int
	DbUser     string
	DbPassword string
	DbName     string
	DbSslMode  string
}

var configurations *Config

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

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("Env error: JWT_SECRET is missing")
	}

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatal("Env error: DB_HOST is missing")
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Fatal("Env error: DB_PORT is missing")
	}

	actualDbPort, err := strconv.Atoi(dbPort)
	if err != nil {
		log.Fatal("Env error: Invalid DB_PORT given\nError: ", err)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal("Env error: DB_USER is missing")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("Env error: DB_PASSWORD is missing")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("Env error: DB_NAME is missing")
	}

	dbSslMode := os.Getenv("DB_SSLMODE")
	if dbSslMode == "" {
		log.Fatal("Env error: DB_SSLMODE is missing")
	}

	configurations = &Config{
		Version:   version,
		AppName:   appName,
		HttpPort:  actualPort,
		JwtSecret: jwtSecret,

		// DB
		DbHost:     dbHost,
		DbPort:     actualDbPort,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		DbName:     dbName,
		DbSslMode:  dbSslMode,
	}
}

// Singleton Design Pattern: creates one n share with all
func GetConfig() *Config {
	if configurations == nil {
		loadConfig()
	}

	return configurations
}
