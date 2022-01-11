package main

import (
	"fmt"
	"log"

	"github.com/capcom6/go-tube/metadata/config"
	"github.com/capcom6/go-tube/metadata/internal/metadata"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config := config.GetConfig(".env")
	repository, err := metadata.NewMetadataRepository(config.DbHost, config.DbName)
	if err != nil {
		panic(err)
	}

	defer repository.Disconnect()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Metadata service online")
	})

	video := app.Group("/video")
	metadata.Register(video, repository)

	// app.Get("/video", func(c *fiber.Ctx) error {
	// 	if id := c.Query("id", ""); id != "" {
	// 		video, err := repository.GetById(id)
	// 		if err == mongo.ErrNoDocuments {
	// 			return c.SendStatus(http.StatusNotFound)
	// 		}
	// 		if err != nil {
	// 			fmt.Printf("Error getting video: %v", err)
	// 			return c.SendStatus(http.StatusInternalServerError)
	// 		}
	// 		return c.JSON(video)
	// 	}

	// 	videos, err := repository.SelectAll()
	// 	if err != nil {
	// 		return err
	// 	}

	// 	if len(videos) == 0 {
	// 		return c.JSON([]string{})
	// 	}

	// 	return c.JSON(videos)
	// })

	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.Port)))
}
