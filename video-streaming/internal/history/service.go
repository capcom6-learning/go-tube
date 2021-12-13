package history

import (
	"encoding/json"

	"github.com/streadway/amqp"
)

type Service struct {
	rabbitMQ string
}

type message struct {
	VideoPath string `json:"videoPath"`
}

func NewService(rabbitMQ string) *Service {
	return &Service{
		rabbitMQ: rabbitMQ,
	}
}

func (s *Service) Send(videoPath string) error {

	conn, err := amqp.Dial(s.rabbitMQ)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	if err := ch.ExchangeDeclare("viewed", "fanout", true, false, false, false, nil); err != nil {
		return err
	}

	msg := message{
		VideoPath: videoPath,
	}

	bytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	err = ch.Publish(
		"viewed", // exchange
		"",       // routing key
		false,    // mandatory
		false,    // immediate
		amqp.Publishing{
			Body: bytes,
		})
	if err != nil {
		return err
	}

	return nil
}
