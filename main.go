package main

import (
	"github.com/snipeart007/gin-gorm-template/app"
	"github.com/snipeart007/gin-gorm-template/app/auth"
	"github.com/snipeart007/gin-gorm-template/state"
	"github.com/snipeart007/gin-gorm-template/types"
)

var DB_DSN string = "host=localhost user=" + state.DB_USER + " password=" + state.DB_PASSWORD + " dbname=" + state.DB_DBNAME + " port=" + state.DB_PORT + " sslmode=disable TimeZone=Asia/Kolkata"

var RouteGroups = []types.RouteInitializer{
	auth.Auth,
}

func main() {
	bookClub := app.CreateApp(DB_DSN)
	bookClub.InitRoutes(RouteGroups)

	bookClub.Start(state.PORT)
}
