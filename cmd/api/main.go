package main

import (
	"log"

	"github.com/Jackson-SM/Europa/cmd/internal/config"
	"github.com/Jackson-SM/Europa/cmd/internal/database"
	"github.com/Jackson-SM/Europa/cmd/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	host := "localhost"
	port := "3030"
	router := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	cfg := config.Load()
	db := database.Connect(cfg.DatabaseURL)

	rg := router.Group("/api/v1")
	routes.UserRoutes(rg, db)

	router.Run(host + ":" + port)
}