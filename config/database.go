package config

import (
	"goFiber/main/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func InitDatabase() {
	db, err := gorm.Open(sqlite.Open("bookmark.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database \n", err.Error())
		os.Exit(2)
	}

	db.AutoMigrate(&models.Bookmark{}, &models.User{})

	Database = DbInstance{Db: db}
}
