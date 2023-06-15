package constants

type VideoQuality string

const (
	Q144p VideoQuality = "144p"
	Q240p VideoQuality = "240p"
	Q360p VideoQuality = "360p"
	Q480p VideoQuality = "480p"
	Q720p VideoQuality = "720p"
)

func (v VideoQuality) String() string {
	return string(v)
}

func GetVideoQualities() [5]VideoQuality {
	return [...]VideoQuality{Q144p, Q240p, Q360p, Q480p, Q720p}
}

func GetVideoQualitySet() map[VideoQuality]struct{} {
	videoQualities := GetVideoQualities()
	videoQualitySet := make(map[VideoQuality]struct{})
	for _, videoQuality := range videoQualities {
		videoQualitySet[videoQuality] = struct{}{}
	}
	return videoQualitySet
}

type VideoStatus string

const (
	VideoStatusPending   VideoStatus = "PENDING"
	VideoStatusCompleted VideoStatus = "COMPLETED"
	VideoStatusFailed    VideoStatus = "FAILED"
)

func (v VideoStatus) String() string {
	return string(v)
}

func GetVideoStatuses() [3]VideoStatus {
	return [...]VideoStatus{VideoStatusPending, VideoStatusCompleted, VideoStatusFailed}
}
