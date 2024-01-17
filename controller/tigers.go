package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type TigersController struct{}

func (tc *TigersController) CreateTiger(ctx *gin.Context) {
	fmt.Println(ctx.Get("email"))
	ctx.JSON(200, gin.H{"message": "Create Tiger"})
}
