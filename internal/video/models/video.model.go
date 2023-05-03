package models

type Video struct {
	Email            string `json:"email" bson:"email"`
	UploadedVideoUrl string `json:"uploadedVideoUrl" bson:"uploadedVideoUrl"`
	EnhancedVideoUrl string `json:"enhancedVideoUrl" bson:"enhancedVideoUrl"`
}
