package types

type RequestInterface string

const (
	UI              RequestInterface = "ui"
	API             RequestInterface = "api"
	ChromeExtension RequestInterface = "chromeExtension"
	Whatsapp        RequestInterface = "whatsapp"
	Discord         RequestInterface = "discord"
	ClientLibrary   RequestInterface = "clientLibrary"
)

func (r RequestInterface) String() string {
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
