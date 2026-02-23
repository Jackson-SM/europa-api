package routes

import (
	"github.com/Jackson-SM/Europa/cmd/internal/user"
	"github.com/gin-gonic/gin"
)

func AddRoutes(route *gin.Engine) {
	routeGroups := route.Group("/api/v1")
	{
		user.AddUserRoutes(routeGroups)
	}
}