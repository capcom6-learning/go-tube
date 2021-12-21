package main

import (
	"log"

	"github.com/capcom6/go-tube/recommendations/config"
	"github.com/capcom6/go-tube/recommendations/internal/history"
)

func main() {
	config := config.GetConfig()

	rabbit, err := history.NewRabbit(config.Rabbit)
	if err != nil {
		log.Fatal(err)
	}

	if err := rabbit.Listen(); err != nil {
		log.Fatal(err)
	}
	defer rabbit.Close()
}
