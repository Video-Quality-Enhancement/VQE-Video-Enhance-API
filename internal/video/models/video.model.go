package models

type Video struct {
	Email            string `json:"email" bson:"email"`
	RequestId        string `json:"requestId" bson:"requestId"` // ? THINK ABOUT THIS
	UploadedVideoUrl string `json:"uploadedVideoUrl" bson:"uploadedVideoUrl"`
	EnhancedVideoUrl string `json:"enhancedVideoUrl,omitempty" bson:"enhancedVideoUrl,omitempty"`
}
