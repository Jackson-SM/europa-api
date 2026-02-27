package main

import (
	"fmt"
	"log"

	"github.com/Jackson-SM/Europa/internal/config"
	"github.com/Jackson-SM/Europa/internal/database"
	"github.com/Jackson-SM/Europa/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found")
	}

	cfg := config.Load()
	db := database.Connect(cfg.DatabaseURL)

	if err := database.RunMigrations(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	rg := router.Group("/api/v1")
	routes.UserRoutes(rg, db)

	addr := fmt.Sprintf(":%s", "3030")

	if err := router.Run(addr); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
