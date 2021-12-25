package main

import (
	"fmt"
	"log"

	"github.com/capcom6/go-tube/metadata/config"
	"github.com/capcom6/go-tube/metadata/internal"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.GetConfig(".env")
	repository, err := internal.NewMetadataRepository(config.DbHost, config.DbName)
	if err != nil {
		panic(err)
	}

	defer repository.Disconnect()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Metadata service online")
	})

	app.Get("/videos", func(c *fiber.Ctx) error {
		videos, err := repository.SelectMetadata()
		if err != nil {
			return err
		}

		return c.JSON(videos)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.Port)))
}
