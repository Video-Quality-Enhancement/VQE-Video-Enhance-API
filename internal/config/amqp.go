package config

import (
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type AMQPconnection interface {
	NewChannel() *amqp.Channel
	DisconnectAll()
}

type amqpConnection struct {
	conn     *amqp.Connection
	channels []*amqp.Channel
}

func NewAMQPconnection() AMQPconnection {

	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	if err != nil {
		slog.Error("Failed to connect to RabbitMQ", "err", err)
		panic(err)
	}

	return &amqpConnection{conn, []*amqp.Channel{}}

}

func (a *amqpConnection) NewChannel() *amqp.Channel {

	ch, err := a.conn.Channel()
	if err != nil {
		slog.Error("Failed to open a channel", "err", err)
		panic(err)
	}

	a.channels = append(a.channels, ch)
	return ch

}

func (a *amqpConnection) DisconnectAll() {

	for _, ch := range a.channels {
		err := ch.Close()
		if err != nil {
			slog.Error("Failed to close a channel", "err", err)
			// not panicing or returning an error
		}
	}

	err := a.conn.Close()
	if err != nil {
		slog.Error("Failed to disconnet from RabbitMQ", "err", err)
		panic(err)
	}

}
