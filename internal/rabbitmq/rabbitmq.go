package rabbitmq

import (
	"context"
	"time"

	"github.com/rabbitmq/amqp091-go"
	"github.com/yxtiblya/internal/cfg"
)

// create and return the channel
func NewChannel() (*amqp091.Channel, error) {
	// returning the connection to rabbitmq
	conn, err := amqp091.Dial(cfg.GetConfig().RabbitURL)
	if err != nil {
		return nil, err
	}

	// create the channel
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return ch, err
}

// send msg in queue
func SendMsg(body []byte, q *amqp091.Queue, ch *amqp091.Channel) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := ch.PublishWithContext(ctx, "", q.Name, false, false, amqp091.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
	if err != nil {
		return err
	}

	return nil
}
