package history

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type video struct {
	VideoPath string `json:"videoPath"`
}

type Handlers struct {
	repository *HistoryRepository
}

func NewHandlers(repository *HistoryRepository) *Handlers {
	return &Handlers{
		repository: repository,
	}
}

func (h *Handlers) Post(c *fiber.Ctx) error {
	v := new(video)

	if err := c.BodyParser(v); err != nil {
		return err
	}

	fmt.Println(v.VideoPath)

	return h.repository.Insert(v.VideoPath)
}
