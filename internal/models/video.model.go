package models

type VideoEnhance struct {
	UserId             string   `json:"userId" bson:"userId"`
	RequestId          string   `json:"requestId" bson:"requestId"`
	ResponseInterfaces []string `json:"responseInterfaces" bson:"responseInterfaces" binding:"required,are-response-interfaces-valid"`
	VideoUrl           string   `json:"videoUrl" bson:"videoUrl" binding:"required,url,is-video-capturable"`
}

type VideoEnhanceRequest struct {
	UserId    string `json:"userId" bson:"userId"`
	RequestId string `json:"requestId" bson:"requestId"`
	VideoUrl  string `json:"videoUrl" bson:"videoUrl"`
}
