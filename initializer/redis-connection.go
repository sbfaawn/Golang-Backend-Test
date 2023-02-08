package initializers

import (
	"context"
	"fmt"
	"log"
)

var Client *redis.Client
var ctx = context.Background()

func ConnectToRedis(config *Config) {
	Client := redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword,
		DB:       config.RedisDBNum,
	})

	pong, err := Client.Ping(ctx).Result()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Redis Connection is Established", pong)
}
