package database

import (
	"Golang-Backend-Test/config"
	"Golang-Backend-Test/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {
	config := config.GetDBConfigurationString()

	log.Print(config)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&model.Student{})

	return db
}

func GetDBInstance() *gorm.DB {
	return DB
}
