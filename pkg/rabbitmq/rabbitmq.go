package rabbitmq

import (
	"context"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	RabbitMQURL string
	con         *amqp091.Connection
	channels    map[string]*amqp091.Channel
}

func (r *RabbitMQ) Connect() (err error) {
	r.con, err = amqp091.Dial(r.RabbitMQURL)
	return
}

func (r *RabbitMQ) Close() (err error) {
	err = r.con.Close()
	return
}

func (r *RabbitMQ) CreateChannel(name string) (err error) {
	r.channels[name], err = r.con.Channel()
	if err != nil {
		return
	}
	_, err = r.channels[name].QueueDeclare(name, false, false, false, false, nil)
	return
}

func (r *RabbitMQ) CloseChannel(name string) (err error) {
	_, err = r.channels[name].QueueDelete(name, false, false, false)
	if err != nil {
		return
	}

	err = r.channels[name].Close()
	return
}

func (r *RabbitMQ) PublishText(ctx context.Context, channelName, body string) (err error) {
	err = r.channels[channelName].PublishWithContext(ctx, "", channelName, false, false, amqp091.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	return
}

func (r *RabbitMQ) Consume(channelName string) (mess <-chan amqp091.Delivery, err error) {
	mess, err = r.channels[channelName].Consume(channelName, "", true, false, false, false, nil)
	return
}

func NewRabbitMQ(rabbitMQURL string) *RabbitMQ {
	return &RabbitMQ{
		RabbitMQURL: rabbitMQURL,
		channels:    make(map[string]*amqp091.Channel),
	}
}
