package model

import (
	"log"


	"go-postgresql/config"
	"fmt"
)

var DB *gorm.DB
var err error

func init() {
	// Connect to DB
	conf := config.Config
	dsn := "host=" + conf.DbHost + " user=" + conf.DbUser + " password=" + conf.DbPassword + " dbname=" + conf.DbName + " port=" + conf.DbPort + " sslmode=disable TimeZone=Asia/Tokyo"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to get result", err)
	}

	// Drop Table if exists
	DB.Exec(`DROP TABLE IF EXISTS posts;`)

	// Make migration
	DB.AutoMigrate(&Post{})
	fmt.Println("migrated")
}
