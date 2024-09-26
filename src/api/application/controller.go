package application

import "github.com/gin-gonic/gin"

func MapRoutes(router *gin.Engine) {
	router.GET("/ping", Ping)
}

func Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}
