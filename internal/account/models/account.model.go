package models

type Account struct {
	Email   string `json:"email" bson:"email"`
	Name    string `json:"name" bson:"name"`
	Picture string `json:"picture" bson:"picture"`
	// WhatsappNumber string `json:"whatsappNumber" bson:"whatsappNumber"`
	// DiscordId string `json:"discordId" bson:"discordId"` // TODO: feature to be added later
}
