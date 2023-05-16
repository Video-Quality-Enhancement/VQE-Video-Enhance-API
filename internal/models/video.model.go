package models

type VideoEnhance struct {
	UserId             string   `json:"userId" bson:"userId"`
	RequestId          string   `json:"requestId" bson:"requestId"`
	ResponseInterfaces []string `json:"responseInterfaces" bson:"responseInterfaces"`
	UploadedVideoUri   string   `json:"UploadedVideoUri" bson:"UploadedVideoUri"`
	EnhancedVideoUri   string   `json:"EnhancedVideoUri,omitempty" bson:"EnhancedVideoUri,omitempty"`
}

type VideoEnhanceRequest struct {
	UserId           string `json:"userId" bson:"userId"`
	RequestId        string `json:"requestId" bson:"requestId"`
	UploadedVideoUri string `json:"UploadedVideoUri" bson:"UploadedVideoUri"`
}

// TODO: change uri back to url and apply binding
