package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/capcom6/go-tube/video-streaming/config"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VideoRecord struct {
	ID   primitive.ObjectID `bson:"_id"`
	Path string             `bson:"videoPath"`
}

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

	dbhost := config.Config("DBHOST")
	dbname := config.Config("DBNAME")
	if dbhost == "" {
		panic("DBHOST enviroment variable not set")
	}
	if dbname == "" {
		panic("DBNAME enviroment variable not set")
	}

	mongodb, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dbhost))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := mongodb.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	collection := mongodb.Database(dbname).Collection("videos")

	app := fiber.New()

	app.Get("/video", func(c *fiber.Ctx) error {
		videoId := c.Query("id")
		if videoId == "" {
			return c.SendStatus(404)
		}
		objectId, err := primitive.ObjectIDFromHex(videoId)
		if err != nil {
			log.Println("Invalid id")
		}

		var video VideoRecord
		if err := collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&video); err != nil {
			return err
		}

		url := fmt.Sprintf("http://%s:%d/video?path=%s", storageHost, storagePort, video.Path)

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
