package config

import (
	"fmt"
	"goFiber/main/models"
	"gorm.io/gorm"
	"log"
	"os"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func InitDatabase(user string, password string, host string, port string, database string) {

	dsn, _ := fmt.Printf("postgres://%s:%s@%s:%s/%s", user, password, host, port, database)
	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database \n", err.Error())
		os.Exit(2)
	}

	db.AutoMigrate(&models.Bookmark{}, &models.User{})

	Database = DbInstance{Db: db}
}
