package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Ping(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		// You can access the db object here
		// Response and Request types from types are available here
		c.JSON(http.StatusOK, PingPong{
			Message: "pong",
		})
	}
}
