package producers

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/models"
	amqp "github.com/rabbitmq/amqp091-go"
)

type NotificationProducer interface {
	PublishNotification(*models.VideoEnhance)
}

type notificationProducer struct {
	conn config.AMQPconnection
}

func NewNotificationProducer() NotificationProducer {
	return &notificationProducer{
		conn: config.NewAMQPconnection(),
	}
}

func (producer *notificationProducer) PublishNotification(video *models.VideoEnhance) {

	ch, err := producer.conn.NewChannel()
	if err != nil {
		log.Printf("%s: %s", "Failed to open a channel", err)
		return
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"notification",
		"topic",
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

	body, err := json.Marshal(video)
	if err != nil {
		log.Printf("%s: %s", "Failed to marshal video object", err)
		return
	}

	// ? Should I pass the correlation id also
	err = ch.PublishWithContext(
		ctx,
		"notification",
		"notify."+video.RequestInterface,
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
