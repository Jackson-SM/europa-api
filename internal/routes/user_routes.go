package routes

import (
	"github.com/Jackson-SM/Europa/internal/controllers"
	repositories "github.com/Jackson-SM/Europa/internal/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(r *gin.RouterGroup, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(userRepository)

	r.GET("/users/:id", userController.FindById)
	r.POST("/users", userController.Create)
}
