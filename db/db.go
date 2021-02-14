package db

import (
	"main/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

//GetDB ....
func GetDB() *gorm.DB {
	return db
}

//SetupDB ...
func SetupDB() {
	//SQLite
	database, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	//MySQL
	// dsn := "root:@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=true&loc=Local"
	// database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	database.AutoMigrate(&model.User{})
	database.AutoMigrate(&model.Product{})
	database.AutoMigrate(&model.Transaction{})

	db = database
}
