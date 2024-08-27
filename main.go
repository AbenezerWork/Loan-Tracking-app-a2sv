package main

import (
	"loan-tracking/api/router"
	"loan-tracking/bootstrap"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	client, err := bootstrap.InitMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	r := gin.Default()

	router.InitRoutes(r, client)

	if err := r.Run(os.Getenv("ADDRESS")); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
