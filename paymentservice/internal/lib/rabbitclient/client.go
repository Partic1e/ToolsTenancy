package rabbitclient

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type RabbitClient struct {
	conn  *amqp.Connection
	ch    *amqp.Channel
	queue string
}

func NewRabbitClient(url, queue string) (*RabbitClient, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		err = conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	_, err = ch.QueueDeclare(queue, true, false, false, false, nil)
	if err != nil {
		err = ch.Close()
		return nil, fmt.Errorf("failed to declare a queue: %w", err)
	}

	return &RabbitClient{conn: conn, ch: ch, queue: queue}, nil
}

func (r *RabbitClient) Publish(event interface{}) error {
	lines, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	err = r.ch.Publish(
		"",
		r.queue,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        lines,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish event: %w", err)
	}

	return nil
}

func (r *RabbitClient) Close() {
	if err := r.ch.Close(); err != nil {
		log.Printf("failed to close channel: %v", err)
	}

	if err := r.conn.Close(); err != nil {
		log.Printf("failed to close connection: %v", err)
	}
}
