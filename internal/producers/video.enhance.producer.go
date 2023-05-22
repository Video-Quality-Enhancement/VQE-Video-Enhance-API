package producers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/models"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/utils"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type VideoEnhanceProducer interface {
	Publish(video *models.VideoEnhanceRequest) error
}

type videoEnhanceProducer struct {
	ch *amqp.Channel
}

func NewVideoEnhanceProducer(ch *amqp.Channel) VideoEnhanceProducer {
	return &videoEnhanceProducer{ch}
}

func (producer *videoEnhanceProducer) Publish(request *models.VideoEnhanceRequest) error {

	exchange := "video.enhance"
	err := producer.ch.ExchangeDeclare(
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
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	quality, err := utils.IdentifyQuality(request.VideoUrl)
	if err != nil {
		slog.Error("Failed to identify quality", "err", err)
		return err
	}

	body, err := json.Marshal(request)
	if err != nil {
		slog.Error("Failed to marshal video object", "err", err)
		return err
	}

	err = producer.ch.PublishWithContext(
		ctx,
		exchange,
		quality,
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
