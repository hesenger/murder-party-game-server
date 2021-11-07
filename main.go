package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var Router *gin.Engine

func main() {
	Router = gin.Default()
	setupApi()

	Router.Run(":" + getPort())
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		return "5000"
	}
	return port
}

func error(ctx *gin.Context, message string) {
	ctx.JSON(500, gin.H{
		"data": gin.H{
			"message": message,
		},
	})
}

func setupApi() {
	api := Router.Group("/api")
	api.GET("/status", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": gin.H{
				"message": "Server is online",
			},
		})
	})

	api.POST("/identify", identify)
}

type IdentifyRequest struct {
	Username string `json:"username" binding:"required"`
}

func identify(ctx *gin.Context) {
	var request IdentifyRequest
	if ctx.BindJSON(&request) != nil {
		error(ctx, "Invalid request body")
		return
	}

	ctx.JSON(200, gin.H{
		"data": gin.H{
			"deviceId": uuid.NewString(),
			"username": request.Username,
		},
	})
}
