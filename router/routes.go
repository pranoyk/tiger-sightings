package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pranoyk/tiger-sightings/controller"
	"github.com/pranoyk/tiger-sightings/model"
	"github.com/pranoyk/tiger-sightings/service"
)

func Init() *gin.Engine {
	router := gin.Default()

	signUpController := controller.SignUpUserController{
		User: &model.SignUpUserRequest{},
		Service: service.NewSignUpUser(),
	}
	loginController := controller.LoginUserController{
		User: &model.LoginRequest{},
		Service: service.NewLogin(),
	}

	router.POST("/register", signUpController.RegisterUser)
	router.POST("/login", loginController.Login)

	return router
}
