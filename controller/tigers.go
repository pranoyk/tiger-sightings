package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pranoyk/tiger-sightings/model"
	"github.com/pranoyk/tiger-sightings/service"
)

type TigersController struct {
	Tiger   *model.CreateTigerRequest
	Service service.Tiger
}

func (tc *TigersController) CreateTiger(ctx *gin.Context) {
	if err := ctx.ShouldBindJSON(tc.Tiger); err != nil {
		fmt.Printf("error binding json: %+v\n", err)
		ctx.JSON(400, gin.H{"error": "Missing fields"})
		return
	}

	email, ok := ctx.Get("email")
	if !ok {
		fmt.Println("error getting email from context")
		ctx.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	err := tc.Service.CreateTiger(ctx.Request.Context(), tc.Tiger, email.(string))
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}
	ctx.JSON(200, gin.H{"message": "Create Tiger"})
}

func (tc *TigersController) CreateSighting(ctx *gin.Context) {
	var sighting model.CreateTigerSightingRequest
	if err := ctx.ShouldBindJSON(&sighting); err != nil {
		fmt.Printf("error binding json: %+v\n", err)
		ctx.JSON(400, gin.H{"error": "Missing fields"})
		return
	}

	sighting.TigerId = ctx.Param("id")
	fmt.Printf("sighting: %+v\n", sighting)

	email, ok := ctx.Get("email")
	if !ok {
		fmt.Println("error getting email from context")
		ctx.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	err := tc.Service.CreateSighting(ctx.Request.Context(), &sighting, email.(string))
	if err != nil {
		ctx.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}
	ctx.JSON(200, gin.H{"message": "Create Sighting"})
}