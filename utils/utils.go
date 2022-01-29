package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snipeart007/gin-gorm-template/state"
)

var log = state.GetLogger()

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func LogError(err error, c *gin.Context) {
	if err != nil {
		log.Print(err.Error())
	}
}
func LogErrorAndRespond(err error, c *gin.Context) {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Print(err.Error())
	}
}
