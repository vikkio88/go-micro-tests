package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct
type Config struct{}

// Get get a config from env
func (c *Config) Get(key string) string {
	return os.Getenv(key)
}

// New returns a new instance of config
func New() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("could not load .env file")
	}
	return Config{}
}
