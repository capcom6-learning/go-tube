package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Rabbit string
}

var instance *Config

func newConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
	}

	rabbit := os.Getenv("RABBIT")
	if rabbit == "" {
		panic("Specify RABBIT enviroment variable")
	}

	return &Config{
		Rabbit: rabbit,
	}
}

func GetConfig() Config {
	if instance == nil {
		instance = newConfig()
	}

	return *instance
}
