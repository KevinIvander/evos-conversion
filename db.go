package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/jinzhu/gorm"
)

var dbConn *gorm.DB
var dbOnce sync.Once

func initDB() {
	dbOnce.Do(func() {
		connString := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_PORT"),
			os.Getenv("MYSQL_DB"),
		)

		conn, err := gorm.Open("mysql", connString)
		if err != nil {
			panic(err)
		}

		dbConn = conn
	})
}
