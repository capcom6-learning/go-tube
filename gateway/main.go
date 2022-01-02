package main

import (
	"fmt"
	"log"

	"github.com/capcom6/go-tube/gateway/config"
	"github.com/capcom6/go-tube/gateway/internal/gateway"
	"github.com/capcom6/go-tube/gateway/internal/metadata"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	config := config.GetConfig(".env")
	metadataService := metadata.NewMetadataService(config.MetadataHost)

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	gateway.Register(app, metadataService)

	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Gateway service online")
	// })

	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.Port)))
}
