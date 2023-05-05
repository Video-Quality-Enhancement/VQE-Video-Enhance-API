package config

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type AMQPconnection interface {
	NewChannel() (*amqp.Channel, error)
	Disconnect()
}

type amqpConnection struct {
	conn *amqp.Connection
}

func NewAMQPconnection() AMQPconnection {

	conn, err := amqp.Dial(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	return &amqpConnection{conn}

}

func (a *amqpConnection) NewChannel() (*amqp.Channel, error) {

	return a.conn.Channel()

}

func (a *amqpConnection) Disconnect() {

	err := a.conn.Close()
	if err != nil {
		log.Panicf("%s: %s", "Failed to disconnet from RabbitMQ", err)
	}

}
