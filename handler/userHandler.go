package handler

import (
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"time"
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

func (us *UserService) ChangePassword(passwordChangeDto *dto.PasswordChangeDto) error {
	if err := validator.New().Struct(passwordChangeDto); err != nil {
		return err
	}

	//Check the user exist with password
	user, err := us.MongoService.LoginCheck(&dto.UserLoginDto{
		Email:    passwordChangeDto.Email,
		Password: passwordChangeDto.OldPassword,
	})
	if err != nil {
		return err
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordChangeDto.NewPassword), 14)
	if err != nil {
		return err
	}
	user.PasswordChanged = time.Now().UTC().Unix()
	user.Password = string(hashedPassword)
	if err := us.MongoService.Update(user); err != nil {
		return err
	}
	return nil
}
