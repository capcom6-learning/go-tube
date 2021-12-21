package history

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type Rabbit struct {
	connection *amqp.Connection
}

type message struct {
	VideoPath string `json:"videoPath"`
}

func NewRabbit(rabbitMQ string) (*Rabbit, error) {
	conn, err := amqp.Dial(rabbitMQ)
	if err != nil {
		return nil, err
	}

	return &Rabbit{
		connection: conn,
	}, nil
}

func (r *Rabbit) Listen() error {

	// defer conn.Close()

	ch, err := r.connection.Channel()
	if err != nil {
		return err
	}
	// defer ch.Close()

	if err := ch.ExchangeDeclare("viewed", "fanout", true, false, false, false, nil); err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	if err := ch.QueueBind(
		q.Name,   // queue name
		"",       // routing key
		"viewed", // exchange
		false,
		nil,
	); err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)

	go func() {
		defer ch.Close()
		for d := range msgs {
			msg := message{}

			if err := json.Unmarshal(d.Body, &msg); err != nil {
				log.Fatal(err)
				continue
			}

			log.Printf("%v\r\n", msg)

			d.Acknowledger.Ack(d.DeliveryTag, false)
		}
		forever <- true
	}()

	<-forever

	return nil
}

func (r *Rabbit) Close() {
	r.connection.Close()
}
