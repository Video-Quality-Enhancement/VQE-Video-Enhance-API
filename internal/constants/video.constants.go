package constants

type ResponseInterface string

const (
	UI              ResponseInterface = "ui"
	API             ResponseInterface = "api"
	ChromeExtension ResponseInterface = "chromeExtension"
	Whatsapp        ResponseInterface = "whatsapp"
	Discord         ResponseInterface = "discord"
	ClientLibrary   ResponseInterface = "clientLibrary"
)

func (r ResponseInterface) String() string {
	return string(r)
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
