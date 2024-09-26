package server

import (
	"github.com/gin-gonic/gin"
	"rate-limiter/src/api/application/controller"
)

func mapRoutes() *gin.Engine {
	router := gin.Default()
	mapControllerRoutes(router)
	return router
}

func mapControllerRoutes(router *gin.Engine) {
	controller.MapRoutes(router)
}
