package handler

import "errors"

func (us *UserService) ValidateUserToken(userId string) error {
	user, err := us.MongoService.FindById(userId)
	if err != nil {
		return err
	}
	if user.PasswordChanged == true {
		return errors.New("invalid token")
	}
	return nil
}
