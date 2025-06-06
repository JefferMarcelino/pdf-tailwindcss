package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BrowserlessHost  string
	BrowserlessToken string
	Port             string
}

func Load() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, falling back to env vars")
	}

	return &Config{
		BrowserlessHost:  os.Getenv("BROWSERLESS_HOST"),
		BrowserlessToken: os.Getenv("BROWSERLESS_TOKEN"),
		Port:             os.Getenv("PORT"),
	}
}
