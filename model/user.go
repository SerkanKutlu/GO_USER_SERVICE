package model

type User struct {
	Id              string `json:"_id" bson:"_id"`
	Name            string `json:"name" bson:"name"`
	Surname         string `json:"surname" bson:"surname"`
	Email           string `json:"email" bson:"email"`
	Password        string `json:"password" bson:"password"`
	Role            string `bson:"role"`
	PasswordChanged bool   `bson:"passwordChanged"`
}
