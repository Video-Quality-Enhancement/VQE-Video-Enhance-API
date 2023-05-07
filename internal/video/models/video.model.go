package models

type VideoEnhance struct {
	Email            string `json:"email" bson:"email"`
	RequestId        string `json:"requestId" bson:"requestId"`
	RequestInterface string `json:"requestInterface" bson:"requestInterface"`
	UploadedVideoUri string `json:"UploadedVideoUri" bson:"UploadedVideoUri"`
	EnhancedVideoUri string `json:"EnhancedVideoUri,omitempty" bson:"EnhancedVideoUri,omitempty"`
}
