package models

type Video struct {
	Email            string `json:"email" bson:"email"`
	CorrelationId    string `json:"correlationId" bson:"correlationId"`
	UploadedVideoUrl string `json:"uploadedVideoUrl" bson:"uploadedVideoUrl"`
	EnhancedVideoUrl string `json:"enhancedVideoUrl,omitempty" bson:"enhancedVideoUrl,omitempty"`
}
