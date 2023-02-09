package main

import (
	"Golang-Backend-Test/entity"
	initializers "Golang-Backend-Test/initializer"
	"fmt"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")

	if err != nil {
		log.Fatalln("? Could not load environment variables", err)
	}

	initializers.ConnectToDB(&config)
}

func main() {
	initializers.DB.AutoMigrate(&entity.OrderItem{}, &entity.User{}, &entity.OrderHistory{})
	fmt.Println("Migration Complete")
}
