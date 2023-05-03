package models

type User struct {
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	Image string `json:"image" bson:"image"`
	// Password string `json:"password" bson:"password"`
}
