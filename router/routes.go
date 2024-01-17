package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/pranoyk/tiger-sightings/controller"
	"github.com/pranoyk/tiger-sightings/middleware"
	"github.com/pranoyk/tiger-sightings/model"
	"github.com/pranoyk/tiger-sightings/repository"
	"github.com/pranoyk/tiger-sightings/service"
)

func Init(db *sql.DB) *gin.Engine {
	router := gin.Default()

	usersRepository := repository.NewUsersRepository(db)
	signUpController := controller.SignUpUserController{
		User:    &model.SignUpUserRequest{},
		Service: service.NewSignUpUser(usersRepository),
	}
	loginController := controller.LoginUserController{
		User:    &model.LoginRequest{},
		Service: service.NewLogin(),
	}
	tigerController := controller.TigersController{
		Tiger:   &model.CreateTigerRequest{},
		Service: service.NewTiger(repository.NewTigersRepository(db)),
	}

	router.POST("/register", signUpController.RegisterUser)
	router.POST("/login", loginController.Login)

	r1 := router.Group("/api/v1")
	r1.Use(middleware.JwtAuthMiddleware())
	r1.POST("/tigers", tigerController.CreateTiger)

	return router
}
