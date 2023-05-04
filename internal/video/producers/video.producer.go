package producers

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

type VideoProducer interface {
	PublishVideo(*models.Video)
	Disconnect()
}

type videoProducer struct {
	conn *amqp.Connection
}

func NewVideoProducer() VideoProducer {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	return &videoProducer{conn}

}

func (producer *videoProducer) Disconnect() {
	err := producer.conn.Close()
	if err != nil {
		log.Panicf("%s: %s", "Failed to disconnet from RabbitMQ", err)
	}
}

func (producer *videoProducer) PublishVideo(video *models.Video) {

	// * creating a new channel for each publish so that we can run this function in a goroutine
	ch, err := producer.conn.Channel()
	if err != nil {
		log.Printf("%s: %s", "Failed to open a channel", err) // ! not sure if this is the right way to handle this
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"video",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("%s: %s", "Failed to declare an exchange", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// quality := indentifyQuality(video) // TODO: need to implement this function
	quality := "144p"

	body, err := json.Marshal(video)
	if err != nil {
		log.Printf("%s: %s", "Failed to marshal video object", err)
		return
	}

	err = ch.PublishWithContext(
		ctx,
		"video",
		quality,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("%s: %s", "Failed to publish a message", err)
		return
	}

	log.Printf("Message Published %s", body)

}
