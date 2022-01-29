package types

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouteInitializer func(engine *gin.Engine, db *gorm.DB)
