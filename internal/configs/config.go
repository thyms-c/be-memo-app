package configs

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI       string `env:"MONGO_URI"`
	MongoDatabase  string `env:"MONGO_DATABASE"`
	Port           string `env:"PORT"`
	UserToken      string `env:"USER_TOKEN"`
	AdminToken     string `env:"ADMIN_TOKEN"`
	FrontendOrigin string `env:"FRONTEND_ORIGIN"`
}

func NewConfig() *Config {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading it. Falling back to system environment variables.")
	}

	config := &Config{}

	// Parse environment variables into the config struct
	if err := env.Parse(config); err != nil {
		log.Fatalln("Failed to parse environment variables into Config struct:", err)
	}

	return config
}
