package models

type VideoEnhance struct {
	UserId           string `json:"userId" bson:"userId"`
	RequestId        string `json:"requestId" bson:"requestId"`
	RequestInterface string `json:"requestInterface" bson:"requestInterface"`
	UploadedVideoUri string `json:"UploadedVideoUri" bson:"UploadedVideoUri"`
	EnhancedVideoUri string `json:"EnhancedVideoUri,omitempty" bson:"EnhancedVideoUri,omitempty"`
}

type VideoEnhanceRequest struct {
	RequestId        string `json:"requestId" bson:"requestId"`
	UploadedVideoUri string `json:"UploadedVideoUri" bson:"UploadedVideoUri"`
}

type VideoEnhanceResponse struct {
	RequestId        string `json:"requestId" bson:"requestId"`
	EnhancedVideoUri string `json:"EnhancedVideoUri" bson:"EnhancedVideoUri"`
}

// TODO: change uri back to url and apply binding
