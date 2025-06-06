package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	BrowserlessHost    string
	BrowserlessToken   string
	TailwindCSSFileURL string
	Port               string
}

func Load() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found, falling back to env vars")
	}

	return &Config{
		BrowserlessHost:    os.Getenv("BROWSERLESS_HOST"),
		BrowserlessToken:   os.Getenv("BROWSERLESS_TOKEN"),
		TailwindCSSFileURL: os.Getenv("TAILWINDCSS_URL"),
		Port:               os.Getenv("PORT"),
	}
}
