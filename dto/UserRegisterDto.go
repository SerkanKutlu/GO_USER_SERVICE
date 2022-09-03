package dto

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"userService/model"
)

type UserRegisterDto struct {
	Name     string `validate:"required"`
	Surname  string `validate:"required"`
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

func (urd *UserRegisterDto) ToUser() (*model.User, error) {

	id := uuid.NewV4().String()
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(urd.Password), 14)
	if err != nil {
		return nil, err
	}
	return &model.User{
		Id:              id,
		Name:            urd.Name,
		Surname:         urd.Surname,
		Email:           urd.Email,
		Password:        string(hashedPassword),
		Role:            "User",
		PasswordChanged: false,
	}, nil
}
