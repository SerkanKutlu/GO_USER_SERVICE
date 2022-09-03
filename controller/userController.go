package controller

import (
	"github.com/labstack/echo/v4"
	"userService/dto"
	"userService/handler"
)

func GetUserController(userService *handler.UserService) *Controller {
	return &Controller{UserService: userService}
}
func (controller *Controller) Login(c echo.Context) error {
	var loginDto dto.UserLoginDto
	//Binding
	if err := c.Bind(&loginDto); err != nil {
		return c.JSON(400, "Invalid Payload")
	}

	token, err := controller.UserService.Login(&loginDto)
	if err != nil {
		return c.JSON(401, "UnAuthorized")
	}
	return c.JSON(200, token)
}

func (controller *Controller) Register(c echo.Context) error {
	var registerDto dto.UserRegisterDto
	//Binding
	if err := c.Bind(&registerDto); err != nil {
		return c.JSON(400, "Invalid Payload")
	}
	createdId, err := controller.UserService.Register(&registerDto)
	if err != nil {
		return c.JSON(404, err.Error())
	}
	return c.JSON(200, createdId)
}
