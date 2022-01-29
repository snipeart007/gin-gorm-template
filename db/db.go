package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/snipeart007/gin-gorm-template/db/models"
	"github.com/snipeart007/gin-gorm-template/utils"
)


func InitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	utils.CheckError(err)

	db.AutoMigrate(&models.User{})

	return db
}