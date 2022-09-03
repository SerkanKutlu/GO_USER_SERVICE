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

func (controller *Controller) ChangePassword(c echo.Context) error {
	var passwordChangeDto dto.PasswordChangeDto
	//Binding
	if err := c.Bind(&passwordChangeDto); err != nil {
		return c.JSON(400, "Invalid Payload")
	}
	if err := controller.UserService.ChangePassword(&passwordChangeDto); err != nil {
		return c.JSON(400, err.Error())
	}
	return c.JSON(200, "changed, have to login again")
}
func (controller *Controller) ValidateToken(c echo.Context) error {
	id := c.Param("id")
	if err := controller.UserService.ValidateUserToken(id); err != nil {
		return c.JSON(401, "not a valid token")
	}
	return c.JSON(200, "ok")
}
