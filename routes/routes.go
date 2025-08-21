package routes

import (
	"learn-golang-gin/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("api/register", controllers.Register)

	return router
}
