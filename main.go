package main

import (
	"learn-golang-gin/config"
	"learn-golang-gin/database"
	"learn-golang-gin/routes"
)

func main() {

	// load Config
	config.LoadEnv()

	// inisialisasi database
	database.InitDB()

	// Setup Routes
	r := routes.SetupRoutes()

	// Sever Run
	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
