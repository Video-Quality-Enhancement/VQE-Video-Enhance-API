package models

import "time"

type VideoEnhance struct {
	UserId               string    `json:"userId" bson:"userId"`
	RequestId            string    `json:"requestId" bson:"requestId"`
	ResponseInterfaces   []string  `json:"responseInterfaces" bson:"responseInterfaces" binding:"required,are-response-interfaces-valid"`
	VideoUrl             string    `json:"videoUrl" bson:"videoUrl" binding:"required,url"`
	VideoQuality         string    `json:"videoQuality" bson:"videoQuality"`
	Status               string    `json:"status" bson:"status"`
	StatusMessage        string    `json:"statusMessage" bson:"statusMessage"`
	EnhancedVideoUrl     string    `json:"enhancedVideoUrl,omitempty" bson:"enhancedVideoUrl,omitempty"`
	EnhancedVideoQuality string    `json:"enhancedVideoQuality,omitempty" bson:"enhancedVideoQuality,omitempty"`
	CreatedAt            time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt" bson:"updatedAt"`
}

type VideoEnhanceRequest struct {
	UserId       string `json:"userId" bson:"userId"`
	RequestId    string `json:"requestId" bson:"requestId"`
	VideoUrl     string `json:"videoUrl" bson:"videoUrl"`
	VideoQuality string `json:"videoQuality" bson:"videoQuality"`
}
