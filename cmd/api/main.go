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

	routes.AddRoutes(router)

	cfg := config.Load()
	database.Connect(cfg.DatabaseURL)

	router.Run(host + ":" + port)
}