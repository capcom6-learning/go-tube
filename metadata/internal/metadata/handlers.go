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

package metadata

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type MetadataHandler struct {
	repository *MetadataRepository
}

func NewMetadataHandler(repository *MetadataRepository) *MetadataHandler {
	return &MetadataHandler{
		repository: repository,
	}
}

func (h *MetadataHandler) get(c *fiber.Ctx) error {
	id := c.Query("id")
	if id == "" {
		return h.selectAll(c)
	}

	video, err := h.repository.GetById(id)
	if err == mongo.ErrNoDocuments {
		return c.SendStatus(fiber.StatusNotFound)
	}
	if err != nil {
		fmt.Printf("Error getting video: %v", err)
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.JSON(video)
}

func (h *MetadataHandler) selectAll(c *fiber.Ctx) error {
	videos, err := h.repository.SelectAll()
	if err != nil {
		return err
	}

	if len(videos) == 0 {
		return c.JSON([]string{})
	}

	return c.JSON(videos)
}

func (h *MetadataHandler) post(c *fiber.Ctx) error {
	input := new(Metadata)

	if err := c.BodyParser(input); err != nil {
		return err
	}

	if input.Name == "" || input.Path == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	_, err := h.repository.Create(input)
	if err != nil {
		return err
	}

	fmt.Printf("%v\r\n", *input)

	return c.JSON(input)
}

func Register(router fiber.Router, repository *MetadataRepository) {
	handler := NewMetadataHandler(repository)

	router.Get("/", handler.get)
	router.Post("/", handler.post)
}
