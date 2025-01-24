package constants

type VideoStatus string

const (
	VideoStatusPending    VideoStatus = "PENDING"
	VideoStatusProcessing VideoStatus = "PROCESSING"
	VideoStatusSuccess    VideoStatus = "SUCCESS"
	VideoStatusFailed     VideoStatus = "FAILED"
)

func (v VideoStatus) String() string {
	return string(v)
}

func GetVideoStatuses() [4]VideoStatus {
	return [...]VideoStatus{VideoStatusPending, VideoStatusProcessing, VideoStatusSuccess, VideoStatusFailed}
}
