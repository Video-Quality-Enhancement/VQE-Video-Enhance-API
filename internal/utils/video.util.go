package utils

import (
	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/constants"
	"gocv.io/x/gocv"
	"golang.org/x/exp/slog"
)

func IdentifyQuality(videoUrl string) (string, error) {

	vc, err := gocv.OpenVideoCapture(videoUrl)
	if err != nil {
		slog.Error("Error opening video capture file using url", "videoUrl", videoUrl)
		return "", err
	}
	defer vc.Close()

	var quality constants.VideoQuality
	height := vc.Get(gocv.VideoCaptureFrameHeight)
	width := vc.Get(gocv.VideoCaptureFrameWidth)
	area := height * width

	// switch {
	// case height <= 144:
	// 	quality = constants.Q144p
	// case height <= 240:
	// 	quality = constants.Q240p
	// case height <= 360:
	// 	quality = constants.Q360p
	// case height <= 480:
	// 	quality = constants.Q480p
	// default:
	// 	quality = constants.Q720p
	// }

	switch {
	case area <= 256*144:
		quality = constants.Q144p
	case area <= 426*240:
		quality = constants.Q240p
	case area <= 640*360:
		quality = constants.Q360p
	case area <= 854*480:
		quality = constants.Q480p
	default:
		quality = constants.Q720p
	}

	slog.Debug("Identified video quality", "quality", quality.String())

	return quality.String(), nil

}
