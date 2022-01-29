package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/snipeart007/gin-gorm-template/db"
	"github.com/snipeart007/gin-gorm-template/types"
)

type App struct {
	Engine *gin.Engine
	DB     *gorm.DB
}

func CreateApp(dsn string) *App {
	engine := gin.Default()
	db := db.InitDB(dsn)

	engine.SetTrustedProxies(nil)

	return &App{
		Engine: engine,
		DB:     db,
	}
}

func (app *App) InitRoutes(routeGroups []types.RouteInitializer) {
	for _, routeGroup := range routeGroups {
		routeGroup(app.Engine, app.DB)
	}
}

func (app *App) Start(port string) {
	app.Engine.Run(":" + port)
}
