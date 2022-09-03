package handler

import (
	"github.com/go-playground/validator/v10"
	"userService/dto"
)

func (us *UserService) Register(registerDto *dto.UserRegisterDto) (*string, error) {
	//Validation
	if err := validator.New().Struct(registerDto); err != nil {
		return nil, err
	}
	//To User
	newUser, err := registerDto.ToUser()
	if err != nil {
		return nil, err
	}

	if err := us.MongoService.Insert(newUser); err != nil {
		return nil, err
	}
	return &newUser.Id, nil

}

func (us *UserService) Login(loginDto *dto.UserLoginDto) (*string, error) {
	//Validation
	if err := validator.New().Struct(loginDto); err != nil {
		return nil, err
	}

	user, err := us.MongoService.LoginCheck(loginDto)
	if err != nil {
		return nil, err
	}
	token, err := us.TokenService.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return token, nil

}
