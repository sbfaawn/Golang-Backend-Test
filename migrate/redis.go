package main

import (
	initializers "Golang-Backend-Test/initializer"
	"context"
	"fmt"
	"log"
)

var ctx = context.Background()

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("? Could not load environment variables", err)
	}

	initializers.ConnectToRedis(&config)
}

func main() {
	err := initializers.Client.Set(initializers.Ctx, "key", "value", 0).Err()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Success")
}
