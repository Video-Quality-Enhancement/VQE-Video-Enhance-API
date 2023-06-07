package producers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/models"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type VideoEnhanceProducer interface {
	Publish(video *models.VideoEnhanceRequest) error
}

type videoEnhanceProducer struct {
	exchange string
	ch       *amqp.Channel
}

func NewVideoEnhanceProducer(ch *amqp.Channel) VideoEnhanceProducer {

	exchange := "video.enhance"
	err := ch.ExchangeDeclare(
		exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		slog.Error("Failed to declare an exchange", "err", err)
		panic(err)
	}

	return &videoEnhanceProducer{exchange, ch}
}

func (producer *videoEnhanceProducer) Publish(request *models.VideoEnhanceRequest) error {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(request)
	if err != nil {
		slog.Error("Failed to marshal video object", "err", err)
		return err
	}

	err = producer.ch.PublishWithContext(
		ctx,
		producer.exchange,
		request.VideoQuality,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		slog.Error("Failed to publish a message", "err", err)
		return err
	}

	slog.Debug("Message Published", "body", body)
	return nil

}
