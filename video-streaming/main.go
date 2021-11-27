package main

import (
	"log"

	"github.com/capcom6/go-tube/config"
	// "github.com/capcom6/go-todo/todo"
	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	port := config.Config("PORT")
	if port == "" {
		panic("PORT enviroment variable not set")
	}

	app := fiber.New()
	// app.Use(cors.New())

	// database.ConnectDB()
	// defer database.Disconnect()

	// api := app.Group("/api")
	// todo.Register(api, database.DB)

	app.Get("/video", func(c *fiber.Ctx) error {
		return c.SendFile("./video/SampleVideo_1280x720_30mb.mp4")
	})

	log.Fatal(app.Listen(":" + config.Config("PORT")))
}
