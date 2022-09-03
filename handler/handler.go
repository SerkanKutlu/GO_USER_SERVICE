package handler

import (
	tokenpkg "userService/internal/token"
	mongodb "userService/repository/mongo"
)

type UserService struct {
	MongoService *mongodb.MongoService
	TokenService *tokenpkg.TokenService
}

var userService = new(UserService)

func GetUserService() *UserService {
	return userService
}
func SetMongoService(mongoService *mongodb.MongoService) {
	userService.MongoService = mongoService
}
func SetTokenService(tokenService *tokenpkg.TokenService) {
	userService.TokenService = tokenService
}
