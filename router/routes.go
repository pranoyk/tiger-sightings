package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pranoyk/tiger-sightings/controller"
	"github.com/pranoyk/tiger-sightings/model"
	"github.com/pranoyk/tiger-sightings/service"
)

func Init() *gin.Engine {
	router := gin.Default()

	controller := controller.UserController{
		User: &model.SignUpUserRequest{},
		Service: service.NewSignUpUser(),
	}

	router.POST("/register", controller.RegisterUser)

	return router
}
