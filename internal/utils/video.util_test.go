package utils_test

import (
	"testing"

	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/utils"
)

func TestIdentifyQuality(t *testing.T) {

	quality, err := utils.IdentifyQuality("https://download.samplelib.com/mp4/sample-5s.mp4")

	t.Log(quality)

	if err != nil {
		t.Error("Failed to identify quality", "error", err)
	}

}
