package producers_test

import (
	"os"
	"testing"

	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/config"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/models"
	"github.com/Video-Quality-Enhancement/VQE-User-Video-API/internal/producers"
)

func TestVideoEnhanceProducer_Publish(t *testing.T) {

	os.Setenv("AMQP_URL", "amqp://guest:guest@localhost:5672/")
	conn := config.NewAMQPconnection()
	defer conn.DisconnectAll()

	producer := producers.NewVideoEnhanceProducer(conn)
	err := producer.Publish(&models.VideoEnhanceRequest{
		UserId:    "1234",
		RequestId: "123",
		VideoUrl:  "https://download.samplelib.com/mp4/sample-5s.mp4",
	})

	if err != nil {
		t.Error("Failed to publish message", "err", err)
	}

}
