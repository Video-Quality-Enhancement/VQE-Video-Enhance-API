package models

import "time"

type Video struct {
	UserId               string `json:"userId" bson:"userId"`
	VideoId              string `json:"videoId" bson:"videoId"`
	VideoResolution      string `json:"videoResolution" bson:"videoResolution"`
	VideoDurationSeconds int    `json:"videoDurationSeconds" bson:"videoDurationSeconds"`
}

type VideoEnhance struct {
	UserId          string    `json:"userId" bson:"userId"`
	RequestId       string    `json:"requestId" bson:"requestId"`
	VideoId         string    `json:"videoId" bson:"videoId"`
	Status          string    `json:"status" bson:"status"`
	EnhancedVideoId string    `json:"enhancedVideoId,omitempty" bson:"enhancedVideoId,omitempty"`
	CreatedAt       time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt" bson:"updatedAt"`
}

type VideoEnhanceRequest struct {
	UserId          string `json:"userId" bson:"userId"`
	RequestId       string `json:"requestId" bson:"requestId"`
	VideoId         string `json:"videoId" bson:"videoId"`
	VideoResolution string `json:"videoResolution" bson:"videoResolution"`
}
