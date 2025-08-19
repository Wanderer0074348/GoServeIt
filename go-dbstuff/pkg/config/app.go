package config

import (
	"fmt"

	"github.com/Wanderer0074348/GoServeIt/go-dbstuff/pkg/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func ConnectDb() {
	gormDb, err := gorm.Open(sqlite.Open("data/books.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = gormDb

	db.AutoMigrate(&models.Book{})
}

func GetDB() *gorm.DB {
	fmt.Println("The Database is connected")
	return db
}
