package main

import (
	"go-task-api/models"
	"go-task-api/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	models.InitDB()

	r := gin.Default()
	routes.RegisterRoutes(r)

	log.Println("🚀 Server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
