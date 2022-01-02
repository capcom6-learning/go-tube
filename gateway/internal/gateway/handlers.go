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

type Handler struct {
	metadataService *metadata.MetadataService
}

func NewHandler(metadataService *metadata.MetadataService) *Handler {
	return &Handler{
		metadataService: metadataService,
	}
}

func (h *Handler) index(c *fiber.Ctx) error {
	videos, err := h.metadataService.SelectMetadata()
	if err != nil {
		return err
	}

	return c.Render("video-list", fiber.Map{
		"Videos": videos,
	})
}

func Register(router fiber.Router, metadataService *metadata.MetadataService) {
	handler := NewHandler(metadataService)

	router.Get("/", handler.index)
}
