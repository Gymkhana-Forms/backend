package main

import (
	"log"
	"os"

	"github.com/Gymkhana-Forms/backend/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	// unauthenticated routes
	routes.AuthRoutes(router)
	routes.TestRouter(router)

	// authenticated routes
	routes.UserRoutes(router)
	routes.AdminRouter(router)

	router.Run(":" + port)
}
