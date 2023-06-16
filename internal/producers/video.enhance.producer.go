package producers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/models"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type VideoEnhanceProducer interface {
	Publish(video *models.VideoEnhanceRequest) error
}

type videoEnhanceProducer struct {
	exchange     string
	exchangeType string
	conn         config.AMQPconnection
}

func NewVideoEnhanceProducer(conn config.AMQPconnection) VideoEnhanceProducer {

	exchange := "video.enhance"
	exchangeType := "direct"
	return &videoEnhanceProducer{exchange, exchangeType, conn}
}

func (producer *videoEnhanceProducer) Publish(request *models.VideoEnhanceRequest) error {

	ch, err := producer.conn.NewChannel()
	if err != nil {
		slog.Error("Failed to open a channel", "error", err)
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		producer.exchange,
		producer.exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		slog.Error("Failed to declare an exchange", "error", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	body, err := json.Marshal(request)
	if err != nil {
		slog.Error("Failed to marshal video object", "error", err)
		return err
	}

	err = ch.PublishWithContext(
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
		slog.Error("Failed to publish a message", "error", err)
		return err
	}

	slog.Debug("Message Published", "body", body)
	return nil

}
