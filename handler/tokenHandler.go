package handler

import "errors"

func (us *UserService) ValidateUserToken(nbf int64, userId string) error {
	user, err := us.MongoService.FindById(userId)
	if err != nil {
		return err
	}
	if int(user.PasswordChanged) > int(nbf) {
		return errors.New("invalid token")
	}
	return nil
}
