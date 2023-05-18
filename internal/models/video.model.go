package models

type VideoEnhance struct {
	UserId             string   `json:"userId" bson:"userId"`
	RequestId          string   `json:"requestId" bson:"requestId"`
	ResponseInterfaces []string `json:"responseInterfaces" bson:"responseInterfaces"`
	UploadedVideoUrl   string   `json:"uploadedVideoUrl" bson:"uploadedVideoUrl"`
}

type VideoEnhanceRequest struct {
	UserId           string `json:"userId" bson:"userId"`
	RequestId        string `json:"requestId" bson:"requestId"`
	UploadedVideoUrl string `json:"uploadedVideoUrl" bson:"uploadedVideoUrl"`
}

// TODO: change uri back to url and apply binding
