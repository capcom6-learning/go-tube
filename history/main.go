package main

import (
	"fmt"
	"log"

	"github.com/capcom6/go-tube/history/config"
	"github.com/capcom6/go-tube/history/internal/history"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.GetConfig()
	repository, err := history.NewHistoryRepository(config.DbHost, config.DbName)
	if err != nil {
		panic(err)
	}

	defer repository.Disconnect()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("History Service Online")
	})

	historyHandlers := history.NewHandlers(repository)
	app.Post("/viewed", historyHandlers.Post)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.Port)))
}
