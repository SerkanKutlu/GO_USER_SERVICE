package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"userService/config"
	"userService/controller"
	"userService/handler"
	tokenpkg "userService/internal/token"
	mongodb "userService/repository/mongo"
)

func main() {
	env := os.Getenv("GO_ENV")

	confManger := config.NewConfigurationManager("yml", "application", env)
	//Mongo
	mongoConf := confManger.GetMongoConfiguration()
	mongoService := mongodb.GetMongoService(mongoConf)
	handler.SetMongoService(mongoService)

	//Token
	TokenConf := confManger.GetJwtConfiguration()
	tokenService := tokenpkg.GetTokenService(TokenConf)
	handler.SetTokenService(tokenService)

	//Service
	userService := handler.GetUserService()
	fmt.Println("NAME")
	fmt.Println(userService.MongoService.Users.Name())
	//Controller
	userController := controller.GetUserController(userService)

	e := echo.New()
	//Order Controls
	e.POST("/api/user/login", userController.Login)
	e.POST("/api/user/register", userController.Register)
	err := e.Start(":9000")
	if err != nil {
		panic(err)
	}

}
