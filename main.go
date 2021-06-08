package main

import (
	"github.com/joho/godotenv"

	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql driver
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	godotenv.Load(".env")
}

func main() {
	initDB()
	initRepositories()
	initServices()
	serveHTTP()
}
