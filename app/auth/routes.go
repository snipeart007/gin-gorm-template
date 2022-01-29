package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Auth(engine *gin.Engine, db *gorm.DB) {
	router := engine.Group("/auth")

	router.GET("/ping", Ping(db))
}
