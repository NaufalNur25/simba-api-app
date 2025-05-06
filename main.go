package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/naufal/simba-api-app/config"
	"github.com/naufal/simba-api-app/database"
	"github.com/naufal/simba-api-app/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()
	database.RunMigration()

	r := gin.Default()
	routes.PaymentMethodRoute(r)
	r.Run(":8000")
}
