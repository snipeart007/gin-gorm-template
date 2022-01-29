package state

import (
	"log"
	"os"
	"sync"
)

var logger *log.Logger
var once sync.Once

func GetLogger() *log.Logger {
    once.Do(func() {
        logger = createLogger("app.log")
    })
    return logger
}

func createLogger(fname string) *log.Logger {
    file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return log.New(file, "[APP] ", log.Lshortfile)
}
