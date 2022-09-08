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
	user, err := ms.FindByEmail(loginDto.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password))
	if err != nil {
		return nil, errors.New("wrong password")
	}
	return user, nil
}

func (ms *MongoService) Update(user *model.User) error {
	result, err := ms.Users.ReplaceOne(context.Background(), bson.M{"_id": user.Id}, user)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (ms *MongoService) FindById(id string) (*model.User, error) {
	foundEntity := ms.Users.FindOne(context.Background(), bson.M{"_id": id})
	if foundEntity.Err() != nil {
		return nil, errors.New("user not found")
	}
	var user *model.User
	err := foundEntity.Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ms *MongoService) FindByEmail(email string) (*model.User, error) {
	foundEntity := ms.Users.FindOne(context.Background(), bson.M{"Email": email})
	if foundEntity.Err() != nil {
		return nil, errors.New("user not found")
	}
	var user *model.User
	err := foundEntity.Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
