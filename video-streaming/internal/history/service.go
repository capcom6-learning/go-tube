package history

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type Service struct {
	host string
	port int16
}

func NewService(host string, port int16) *Service {
	return &Service{
		host: host,
		port: port,
	}
}

func (s *Service) Send(videoPath string) error {

	agent := fiber.Post(fmt.Sprintf("http://%s:%d/viewed", s.host, s.port))
	agent.JSON(fiber.Map{"videoPath": videoPath})

	_, _, err := agent.String()

	if len(err) > 0 {
		return err[0]
	}

	return nil
}
