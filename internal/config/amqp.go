package config

import (
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type AMQPconnection interface {
	NewChannel() (*amqp.Channel, error)
	DisconnectAll()
}

type amqpConnection struct {
	conn *amqp.Connection
}

func NewAMQPconnection() AMQPconnection {

	conn, err := amqp.Dial(os.Getenv("AMQP_URL"))
	if err != nil {
		slog.Error("Failed to connect to RabbitMQ", "err", err)
		panic(err)
	}

	return &amqpConnection{conn}

}

func (a *amqpConnection) NewChannel() (*amqp.Channel, error) {

	return a.conn.Channel()

}

func (a *amqpConnection) DisconnectAll() {

	err := a.conn.Close()
	if err != nil {
		slog.Error("Failed to disconnet from RabbitMQ", "err", err)
		panic(err)
	}

}
