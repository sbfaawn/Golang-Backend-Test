package main

import (
	initializers "Golang-Backend-Test/initializer"
	"context"
	"fmt"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("? Could not load environment variables", err)
	}

	initializers.ConnectToRedis(&config)
}

func main() {
	ctx := context.Background()
	cmd := initializers.Client.Set(ctx, "hello", "world", 0)

	fmt.Println(cmd.Result())
}
