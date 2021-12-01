package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/capcom6/go-tube/video-streaming/config"
	"github.com/gofiber/fiber/v2"
)

func main() {
	port := config.Config("PORT")
	if port == "" {
		panic("PORT enviroment variable not set")
	}

	storageHost := config.Config("VIDEO_STORAGE_HOST")
	storagePort, err := strconv.ParseInt(config.Config("VIDEO_STORAGE_PORT"), 10, 32)
	if err != nil {
		panic(err)
	}

	if storageHost == "" {
		panic("VIDEO_STORAGE_HOST enviroment variable not set")
	}
	if storagePort == 0 {
		panic("VIDEO_STORAGE_PORT enviroment variable not set")
	}

	app := fiber.New()

	app.Get("/video", func(c *fiber.Ctx) error {
		url := fmt.Sprintf("http://%s:%d/video?path=SampleVideo_1280x720_30mb.mp4", storageHost, storagePort)

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

	log.Fatal(app.Listen(":" + config.Config("PORT")))
}
