package app

import (
	"os"

	"github.com/Video-Quality-Enhancement/VQE-Video-Enhance-API/internal/services/gapi"
)

func SetUpStorageCORS() {

	uploadVideoGCS := gapi.NewGoogleCloudStorageCORS(os.Getenv("UPLOAD_VIDEO_BUCKET_NAME"))
	enhancedVideoGCS := gapi.NewGoogleCloudStorageCORS(os.Getenv("ENHANCED_VIDEO_BUCKET_NAME"))

	err := uploadVideoGCS.SetUpCORS()
	if err != nil {
		panic(err)
	}

	err = enhancedVideoGCS.SetUpCORS()
	if err != nil {
		panic(err)
	}

}
