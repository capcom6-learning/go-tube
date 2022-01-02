package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         int16
	MetadataHost string
}

var instance *Config

func newConfig(cfgName string) *Config {
	if cfgName != "" {
		godotenv.Load(cfgName)
	}

	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 16)
	if err != nil {
		panic("Incorrect port in PORT enviroment variable")
	}

	metadataHost := os.Getenv("METADATA_HOST")
	if metadataHost == "" {
		panic("Specify METADATA_HOST enviroment variable")
	}

	return &Config{
		Port:         int16(port),
		MetadataHost: metadataHost,
	}
}

func GetConfig(cfgName string) Config {
	if instance == nil {
		instance = newConfig(cfgName)
	}

	return *instance
}
