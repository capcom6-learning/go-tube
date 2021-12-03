package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/capcom6/go-tube/video-streaming/config"
	"github.com/capcom6/go-tube/video-streaming/internal/data/video"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	config := config.GetConfig()
	repository, err := video.NewVideoRepository(config.DbHost, config.DbName)
	if err != nil {
		panic(err)
	}

	defer repository.Disconnect()

	app := fiber.New()

	app.Get("/video", func(c *fiber.Ctx) error {
		videoId := c.Query("id")
		if videoId == "" {
			return c.SendStatus(404)
		}

		video, err := repository.GetVideo(videoId)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.SendStatus(404)
			}
			return c.SendStatus(500)
		}

		url := fmt.Sprintf("http://%s:%d/video?path=%s", config.VideoStorageHost, config.VideoStoragePort, video.Path)

		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return c.SendStatus(500)
		}

		c.Request().Header.VisitAll(func(key, value []byte) {
			name := string(key)
			if name == "Host" || name == "Connection" {
				return
			}

			req.Header.Add(string(key), string(value))
		})

		resp, err := client.Do(req)
		if err != nil {
			return c.SendStatus(500)
		}

		for name, value := range resp.Header {
			if name == "Connection" {
				continue
			}
			c.Set(name, value[0])
		}

		c.Status(resp.StatusCode)

		length, err := strconv.Atoi(resp.Header["Content-Length"][0])
		if err != nil {
			return c.SendStatus(500)
		}

		return c.SendStream(resp.Body, length)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.Port)))
}
