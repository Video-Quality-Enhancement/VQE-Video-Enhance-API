package producers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/models"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/utils"
	amqp "github.com/rabbitmq/amqp091-go"
	"golang.org/x/exp/slog"
)

type VideoEnhanceProducer interface {
	PublishVideo(*models.VideoEnhanceRequest)
}

type videoEnhanceProducer struct {
	conn config.AMQPconnection
}

func NewVideoEnhanceProducer() VideoEnhanceProducer {
	return &videoEnhanceProducer{
		conn: config.NewAMQPconnection(),
	}
}

func (producer *videoEnhanceProducer) PublishVideo(request *models.VideoEnhanceRequest) {

	// * creating a new channel for each publish so that we can run this function in a goroutine
	ch, err := producer.conn.NewChannel()
	if err != nil {
		slog.Error("%s: %s", "Failed to open a channel", err) // ! not sure if this is the right way to handle this
		return
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
		slog.Error("%s: %s", "Failed to declare an exchange", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	quality, err := utils.IdentifyQuality(request.UploadedVideoUri)
	if err != nil {
		slog.Error("%s: %s", "Failed to identify quality", err)
		return
	}

	body, err := json.Marshal(request)
	if err != nil {
		slog.Error("%s: %s", "Failed to marshal video object", err)
		return
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
		slog.Error("%s: %s", "Failed to publish a message", err)
		return
	}

	slog.Debug("Message Published %s", body)

}
