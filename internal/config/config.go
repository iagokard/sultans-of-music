package config

import (
	"log"
	"os"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
	ServerAddr string
	Timezone   string
	JWTSecret  string
}

func Load() *Config {
	return &Config{
		DBUsername: mustGetEnv("SOM_DB_USERNAME"),
		DBPassword: mustGetEnv("SOM_DB_PASSWORD"),
		DBName:     mustGetEnv("SOM_DB_NAME"),
		DBHost:     mustGetEnv("SOM_DB_HOST"),
		DBPort:     mustGetEnv("SOM_DB_PORT"),
		ServerAddr: getEnv("SOM_ADDR", ":8080"),
		JWTSecret:  mustGetEnv("JWT_SECRET"),
	}
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("environment variable %s is required", key)
	}
	return value
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
