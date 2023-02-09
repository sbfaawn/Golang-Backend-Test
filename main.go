package main

import (
	initializers "Golang-Backend-Test/initializer"
	router "Golang-Backend-Test/router"
	"log"
)

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
		return
	}

	e := router.NewRouter()
	e.Logger.Fatal(e.Start(":" + config.ServerPort))
}
