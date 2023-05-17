package constants

type ResponseInterface string

const (
	UI              ResponseInterface = "ui"
	ChromeExtension ResponseInterface = "chromeExtension"
	Whatsapp        ResponseInterface = "whatsapp"
	Discord         ResponseInterface = "discord"
)

func (r ResponseInterface) String() string {
	return string(r)
}

func GetResponseInterfaces() [4]ResponseInterface {
	return [...]ResponseInterface{UI, ChromeExtension, Whatsapp, Discord}
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
