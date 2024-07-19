package main

import (
	"log"

	"github.com/devops-ferraz/the-cats/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	app := gin.Default()
	routes.AppRoutes(app)
	app.Run("localhost:3001")
}
