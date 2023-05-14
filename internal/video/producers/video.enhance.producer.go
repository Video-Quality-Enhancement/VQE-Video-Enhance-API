package producers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-API-Server/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-API-Server/internal/video/models"
	"github.com/Video-Quality-Enhancement/VQE-API-Server/internal/video/utils"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type VideoEnhanceProducer interface {
	Publish(video *models.VideoEnhanceRequest) error
}

type videoEnhanceProducer struct {
	conn config.AMQPconnection
}

func NewVideoEnhanceProducer(conn config.AMQPconnection) VideoEnhanceProducer {
	return &videoEnhanceProducer{
		conn: conn,
	}
}

func (producer *videoEnhanceProducer) Publish(request *models.VideoEnhanceRequest) error {

	// * creating a new channel for each publish so that we can run this function in a goroutine
	ch, err := producer.conn.NewChannel()
	if err != nil {
		slog.Error("Failed to open a channel", "err", err) // ! not sure if this is the right way to handle this
		return err
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"video.enhance",
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

	quality, err := utils.IdentifyQuality(request.UploadedVideoUri)
	if err != nil {
		slog.Error("Failed to identify quality", "err", err)
		return err
	}

	body, err := json.Marshal(request)
	if err != nil {
		slog.Error("Failed to marshal video object", "err", err)
		return err
	}

	// ? Should I pass the correlation id also
	err = ch.PublishWithContext(
		ctx,
		"video.enhance",
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
