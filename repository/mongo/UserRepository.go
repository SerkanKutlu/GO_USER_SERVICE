package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"userService/dto"
	"userService/model"
)

func (ms *MongoService) Insert(entity *model.User) error {
	_, err := ms.Users.InsertOne(context.Background(), entity)
	if err != nil {
		return err
	}
	return nil
}

func (ms *MongoService) LoginCheck(loginDto *dto.UserLoginDto) (*model.User, error) {
	foundEntity := ms.Users.FindOne(context.Background(), bson.M{"email": loginDto.Email})
	if foundEntity.Err() != nil {
		return nil, errors.New("user not found")
	}
	var user model.User
	err := foundEntity.Decode(&user)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password))
	if err != nil {
		return nil, errors.New("wrong password")
	}
	return &user, nil
}
