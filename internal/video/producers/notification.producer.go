package producers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/models"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type NotificationProducer interface {
	PublishNotification(video *models.VideoEnhance)
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
		slog.Error("Failed to open a channel", "err", err)
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
		slog.Error("Failed to declare an exchange", "err", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(video)
	if err != nil {
		slog.Error("Failed to marshal video object", "err", err)
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
		slog.Error("Failed to publish a message", "err", err)
		return
	}

	slog.Debug("Message Published", "body", body)

}
