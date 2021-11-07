package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

type identifyRequest struct {
	Username string `json:"username" binding:"required"`
}

func identify(ctx *gin.Context) {
	var request identifyRequest
	if ctx.BindJSON(&request) != nil {
		error(ctx, "Invalid request body")
		return
	}

	deviceId := uuid.NewString()
	atClaims := jwt.MapClaims{}
	atClaims["deviceId"] = deviceId
	jwt := createJwt(&atClaims)

	ctx.JSON(200, gin.H{
		"data": gin.H{
			"deviceId": deviceId,
			"username": request.Username,
			"token":    jwt,
		},
	})
}

func createJwt(claims *jwt.MapClaims) string {
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, _ := at.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return token
}
