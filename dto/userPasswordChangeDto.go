package dto

type PasswordChangeDto struct {
	Email       string `json:"email" validate:"required,email"`
	OldPassword string `json:"oldPassword" validate:"required"`
	NewPassword string `json:"newPassword" validate:"required"`
}
