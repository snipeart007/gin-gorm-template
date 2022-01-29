package state

import (
	"os"
)

var (
	DB_USER     string = os.Getenv("DB_USER")
	DB_PASSWORD string = os.Getenv("DB_PASSWORD")
	DB_DBNAME   string = os.Getenv("DB_DBNAME")
	DB_PORT     string = os.Getenv("DB_PORT")
	JWT_SECRET  string = os.Getenv("JWT_SECRET")
	PORT        string = os.Getenv("PORT")
)
