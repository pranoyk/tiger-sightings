package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pranoyk/tiger-sightings/model"
	"github.com/pranoyk/tiger-sightings/service"
)

type LoginUserController struct {
	User    *model.LoginRequest
	Service service.Login
}

func (uc *LoginUserController) Login(ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(uc.User); err != nil {
		ctx.JSON(400, gin.H{"error": "Missing fields"})
		return
	}

	access_token, err := uc.Service.LogIn(ctx.Request.Context() ,uc.User)
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(200, gin.H{"user_id": fmt.Sprintf("%v", access_token)})
}
