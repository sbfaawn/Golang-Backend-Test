package main

import (
	initializers "Golang-Backend-Test/initializer"
	"Golang-Backend-Test/model"
	"fmt"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")

	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectToDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&model.OrderItem{}, &model.User{}, &model.OrderHistory{})
	fmt.Println("Migration Complete")
}
