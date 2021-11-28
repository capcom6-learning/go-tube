package main

import (
	"log"

	"github.com/capcom6/go-tube/video-streaming/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	port := config.Config("PORT")
	if port == "" {
		panic("PORT enviroment variable not set")
	}

	app := fiber.New()

	app.Get("/video", func(c *fiber.Ctx) error {
		return c.SendFile("./video/SampleVideo_1280x720_30mb.mp4")
	})

	log.Fatal(app.Listen(":" + config.Config("PORT")))
}
