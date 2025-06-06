package main

import (
	"learn-golang-gin/config"
	"learn-golang-gin/database"

	"github.com/gin-gonic/gin"
)

func main() {

	// load Config
	config.LoadEnv()

	// inisialisasi database
	database.InitDB()

	// inisialisai Gin
	router := gin.Default()

	// create route with method GET
	router.GET("/", func(c *gin.Context) {

		// return response JSON
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	// Sever Run
	router.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
