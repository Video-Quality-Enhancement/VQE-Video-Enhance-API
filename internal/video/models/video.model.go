package models

type VideoEnhance struct {
	Email            string `json:"email" bson:"email"`
	RequestId        string `json:"requestId" bson:"requestId"`
	RequestInterface string `json:"requestInterface" bson:"requestInterface"`
	UploadedVideoUrl string `json:"uploadedVideoUrl" bson:"uploadedVideoUrl"`
	EnhancedVideoUrl string `json:"enhancedVideoUrl,omitempty" bson:"enhancedVideoUrl,omitempty"`
}
