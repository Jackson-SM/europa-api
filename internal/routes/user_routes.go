package routes

import (
	"github.com/Jackson-SM/Europa/cmd/internal/controllers"
	repositories "github.com/Jackson-SM/Europa/cmd/internal/repositories/gorm"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)

	r.GET("/users", userController.Create)
}