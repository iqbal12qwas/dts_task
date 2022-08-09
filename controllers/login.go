package controllers

import (
	"gin-api/credential"
	"gin-api/service"

	"github.com/gin-gonic/gin"
)

//login contorller interface
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jWtService   service.JWTService
}

func LoginHandler(loginService service.LoginService,
	jWtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jWtService:   jWtService,
	} 
}

func (controller *loginController) Login(c *gin.Context) string {
	var credential credential.LoginCredentials
	err := c.ShouldBind(&credential)
	if err != nil {
		return ""
	}
	isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controller.jWtService.GenerateToken(credential.Email, true)
	}
	return ""
}