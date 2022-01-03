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

package gateway

import (
	"github.com/capcom6/go-tube/gateway/internal/metadata"
	"github.com/gofiber/fiber/v2"
)

type GatewayHandler struct {
	metadataService *metadata.MetadataService
}

func NewGatewayHandler(metadataService *metadata.MetadataService) *GatewayHandler {
	return &GatewayHandler{
		metadataService: metadataService,
	}
}

func (h *GatewayHandler) index(c *fiber.Ctx) error {
	videos, err := h.metadataService.SelectMetadata()
	if err != nil {
		return err
	}

	return c.Render("video-list", fiber.Map{
		"Videos": videos,
	})
}

func (h *GatewayHandler) upload(c *fiber.Ctx) error {
	return c.Render("upload-video", fiber.Map{})
}

func (h *GatewayHandler) history(c *fiber.Ctx) error {
	return c.Render("history", fiber.Map{})
}

func (h *GatewayHandler) video(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Redirect("/")
	}

	return c.Render("play-video", fiber.Map{})
}

func Register(router fiber.Router, metadataService *metadata.MetadataService) {
	handler := NewGatewayHandler(metadataService)

	router.Get("/", handler.index)
	router.Get("/upload", handler.upload)
	router.Get("/history", handler.history)
	router.Get("/video", handler.video)
}
