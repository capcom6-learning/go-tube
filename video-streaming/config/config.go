package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port             int16
	VideoStorageHost string
	VideoStoragePort int16
	MetadataUrl      string
	RabbitMQ         string
}

var instance *Config

func newConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("Error loading .env file")
	}

	port, err := strconv.ParseInt(os.Getenv("PORT"), 10, 16)
	if err != nil {
		panic("Incorrect port in PORT enviroment variable")
	}

	videoStorageHost := os.Getenv("VIDEO_STORAGE_HOST")
	if videoStorageHost == "" {
		panic("Specify VIDEO_STORAGE_HOST enviroment variable")
	}

	videoStoragePort, err := strconv.ParseInt(os.Getenv("VIDEO_STORAGE_PORT"), 10, 16)
	if err != nil {
		panic("Incorrect port in VIDEO_STORAGE_PORT enviroment variable")
	}

	metadataUrl := os.Getenv("METADATA_URL")
	if metadataUrl == "" {
		panic("Specify METADATA_URL enviroment variable")
	}

	rabbitMQ := os.Getenv("RABBIT")
	if rabbitMQ == "" {
		panic("Specify RABBIT enviroment variable")
	}

	return &Config{
		Port:             int16(port),
		VideoStorageHost: videoStorageHost,
		VideoStoragePort: int16(videoStoragePort),
		MetadataUrl:      metadataUrl,
		RabbitMQ:         rabbitMQ,
	}
}

func GetConfig() Config {
	if instance == nil {
		instance = newConfig()
	}

	return *instance
}
