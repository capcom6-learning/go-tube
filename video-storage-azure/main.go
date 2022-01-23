package main

import (
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/capcom6/go-tube/video-storage-azure/config"
	"github.com/capcom6/go-tube/video-storage-azure/internal/storage"
	"github.com/gofiber/fiber/v2"
)

func main() {
	port := config.Config("PORT")
	storageAccountName := config.Config("STORAGE_ACCOUNT_NAME")
	storageAccessKey := config.Config("STORAGE_ACCESS_KEY")
	if port == "" {
		panic("Please set PORT enviroment variable.")
	}
	if storageAccountName == "" {
		panic("Please set STORAGE_ACCOUNT_NAME enviroment variable.")
	}
	if storageAccessKey == "" {
		panic("Please set STORAGE_ACCESS_KEY enviroment variable.")
	}

	cred, err := azblob.NewSharedKeyCredential(storageAccountName, storageAccessKey)
	if err != nil {
		panic(err)
	}
	serviceClient, err := azblob.NewServiceClientWithSharedKey(fmt.Sprintf("https://%s.blob.core.windows.net/", storageAccountName), cred, nil)
	if err != nil {
		panic(err)
	}

	// blockBlob := container.NewBlockBlobClient(filename)

	app := fiber.New()

	container := serviceClient.NewContainerClient("videos")
	storage.Register(app, &container)

	// app.Get("/video", func(c *fiber.Ctx) error {
	// 	filename := c.Query("path")
	// 	if filename == "" {
	// 		return c.SendStatus(404)
	// 	}

	// 	cred, err := azblob.NewSharedKeyCredential(storageAccountName, storageAccessKey)
	// 	if err != nil {
	// 		return c.SendStatus(500)
	// 	}
	// 	serviceClient, err := azblob.NewServiceClientWithSharedKey(fmt.Sprintf("https://%s.blob.core.windows.net/", storageAccountName), cred, nil)
	// 	if err != nil {
	// 		return c.SendStatus(500)
	// 	}

	// 	container := serviceClient.NewContainerClient("videos")
	// 	blockBlob := container.NewBlockBlobClient(filename)

	// 	ctx := context.Background()
	// 	properties, err := blockBlob.GetProperties(ctx, &azblob.GetBlobPropertiesOptions{})
	// 	if err != nil {
	// 		return c.SendStatus(500)
	// 	}

	// 	c.Set("Content-Type", *properties.ContentType)
	// 	c.Set("Accept-Ranges", "bytes")

	// 	length := int64(0)
	// 	downloadOption := azblob.DownloadBlobOptions{}
	// 	bytesRange, err := c.Range(int(*properties.ContentLength))
	// 	if err != nil {
	// 		length = *properties.ContentLength
	// 		c.Status(200)
	// 	} else {
	// 		// support of single range only
	// 		begin, end := int64(bytesRange.Ranges[0].Start), int64(bytesRange.Ranges[0].End)
	// 		length = end - begin + 1

	// 		downloadOption.Offset = &begin
	// 		downloadOption.Count = &length

	// 		c.Status(206)
	// 		c.Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", bytesRange.Ranges[0].Start, bytesRange.Ranges[0].End, *properties.ContentLength))
	// 	}

	// 	get, err := blockBlob.Download(ctx, &downloadOption)
	// 	if err != nil {
	// 		return c.SendStatus(500)
	// 	}
	// 	reader := get.Body(azblob.RetryReaderOptions{})

	// 	return c.SendStream(reader, int(length))
	// })

	log.Fatal(app.Listen(":" + port))
}
