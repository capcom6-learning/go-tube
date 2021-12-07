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
	DbHost           string
	DbName           string
	HistoryHost      string
	HistoryPort      int16
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

	dbHost := os.Getenv("DBHOST")
	if dbHost == "" {
		panic("Specify DBHOST enviroment variable")
	}

	dbName := os.Getenv("DBNAME")
	if dbName == "" {
		panic("Specify DBNAME enviroment variable")
	}

	historyHost := os.Getenv("HISTORY_HOST")
	if videoStorageHost == "" {
		panic("Specify HISTORY_HOST enviroment variable")
	}

	historyPort, err := strconv.ParseInt(os.Getenv("HISTORY_PORT"), 10, 16)
	if err != nil {
		panic("Incorrect port in HISTORY_PORT enviroment variable")
	}

	return &Config{
		Port:             int16(port),
		VideoStorageHost: videoStorageHost,
		VideoStoragePort: int16(videoStoragePort),
		DbHost:           dbHost,
		DbName:           dbName,
		HistoryHost:      historyHost,
		HistoryPort:      int16(historyPort),
	}
}

func GetConfig() Config {
	if instance == nil {
		instance = newConfig()
	}

	return *instance
}
