package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pranoyk/tiger-sightings/model"
	"github.com/pranoyk/tiger-sightings/service"
)

type SignUpUserController struct {
	User    *model.SignUpUserRequest
	Service service.SignUpUser
}

func (uc *SignUpUserController) RegisterUser(ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(uc.User); err != nil {
		ctx.JSON(400, gin.H{"error": "Missing fields"})
		return
	}

	userId, err := uc.Service.SignUp(ctx.Request.Context() ,uc.User)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(200, gin.H{"user_id": fmt.Sprintf("%v", userId)})
}
