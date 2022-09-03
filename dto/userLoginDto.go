package dto

type UserLoginDto struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}
