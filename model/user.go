package model

type User struct {
	Id              string `json:"_id" bson:"_id"`
	Name            string `json:"name" bson:"Name"`
	Surname         string `json:"surname" bson:"Surname"`
	Email           string `json:"email" bson:"Email"`
	Password        string `json:"password" bson:"Password"`
	Role            string `bson:"Role"`
	PasswordChanged int64  `bson:"PasswordChanged"`
}
