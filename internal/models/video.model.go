package models

type VideoEnhance struct {
	UserId             string   `json:"userId" bson:"userId"`
	RequestId          string   `json:"requestId" bson:"requestId"`
	ResponseInterfaces []string `json:"responseInterfaces" bson:"responseInterfaces"`
	UploadedVideoUrl   string   `json:"UploadedVideoUrl" bson:"UploadedVideoUrl"`
}

type VideoEnhanceRequest struct {
	UserId           string `json:"userId" bson:"userId"`
	RequestId        string `json:"requestId" bson:"requestId"`
	UploadedVideoUrl string `json:"UploadedVideoUrl" bson:"UploadedVideoUrl"`
}

// TODO: change uri back to url and apply binding
