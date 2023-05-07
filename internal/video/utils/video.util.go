package utils

import (
	"log"

	"github.com/Video-Quality-Enhancement/VQE-Backend/internal/video/types"
	"gocv.io/x/gocv"
)

func IdentifyQuality(videoUrl string) (string, error) {

	vc, err := gocv.VideoCaptureFile(videoUrl)
	if err != nil {
		log.Println("Error opening video capture device: ", videoUrl)
		return "", err
	}
	defer vc.Close()

	var quality types.VideoQuality
	height := vc.Get(gocv.VideoCaptureFrameHeight)

	switch {
	case height <= 144:
		quality = types.Q144p
	case height <= 240:
		quality = types.Q240p
	case height <= 360:
		quality = types.Q360p
	case height <= 480:
		quality = types.Q480p
	default:
		quality = types.Q720p
	}

	return quality.String(), nil

}
