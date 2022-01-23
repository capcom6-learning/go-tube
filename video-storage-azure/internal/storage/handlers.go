// Copyright 2022 Aleksandr Soloshenko
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/gofiber/fiber/v2"
)

const (
	AzureTimeout = 30
)

type StorageHandler struct {
	container *azblob.ContainerClient
}

func (h *StorageHandler) download(c *fiber.Ctx) error {
	filename := c.Query("path")
	if filename == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	blockBlob := h.container.NewBlockBlobClient(filename)

	ctx := c.Context()
	properties, err := blockBlob.GetProperties(ctx, &azblob.GetBlobPropertiesOptions{})
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	c.Set("Content-Type", *properties.ContentType)
	c.Set("Accept-Ranges", "bytes")

	length := int64(0)
	downloadOption := azblob.DownloadBlobOptions{}
	bytesRange, err := c.Range(int(*properties.ContentLength))
	if err != nil {
		length = *properties.ContentLength
		c.Status(200)
	} else {
		fmt.Println(c.Get("Range"))
		// support of single range only
		begin, end := int64(bytesRange.Ranges[0].Start), int64(bytesRange.Ranges[0].End)
		length = end - begin + 1

		downloadOption.Offset = &begin
		downloadOption.Count = &length

		c.Status(206)
		c.Set("Content-Range", fmt.Sprintf("bytes %d-%d/%d", begin, end, *properties.ContentLength))

		fmt.Println(c.GetRespHeader("Content-Range"))
	}

	get, err := blockBlob.Download(ctx, &downloadOption)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	reader := get.Body(azblob.RetryReaderOptions{})

	return c.SendStream(reader, int(length))
}

func (h *StorageHandler) upload(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotImplemented)
}

func NewStorageHandler(c *azblob.ContainerClient) *StorageHandler {
	return &StorageHandler{
		container: c,
	}
}

func Register(r fiber.Router, c *azblob.ContainerClient) {
	handler := NewStorageHandler(c)

	r.Get("/video", handler.download)
	r.Post("/video", handler.upload)
}
