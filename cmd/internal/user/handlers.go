package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Specific User"})
}