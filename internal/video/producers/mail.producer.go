package producers

import (
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/models"
)

type MailProducer interface {
	PublishMail(*models.Video)
}

type mailProducer struct {
	conn config.AMQPconnection
}

func NewMailProducer() MailProducer {
	return &mailProducer{
		conn: config.NewAMQPconnection(),
	}
}

func (producer *mailProducer) PublishMail(video *models.Video) {
	// TODO: Implement mail publisher
}
