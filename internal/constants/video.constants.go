package constants

type ResponseInterface string

const (
	UI       ResponseInterface = "ui"
	Email    ResponseInterface = "email"
	Whatsapp ResponseInterface = "whatsapp"
	Discord  ResponseInterface = "discord"
)

func (r ResponseInterface) String() string {
	return string(r)
}

func GetResponseInterfaces() [4]ResponseInterface {
	return [...]ResponseInterface{Email, UI, Whatsapp, Discord}
}

func GetResponseInterfaceSet() map[ResponseInterface]struct{} {
	responseInterfaces := GetResponseInterfaces()
	responseInterfacesSet := make(map[ResponseInterface]struct{})
	for _, responseInterface := range responseInterfaces {
		responseInterfacesSet[responseInterface] = struct{}{}
	}
	return responseInterfacesSet
}

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
