package initializers

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
)

var Client *redis.Client
var Ctx = context.Background()

func ConnectToRedis(config *Config) {
	Client = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword,
		DB:       config.RedisDBNum,
	})

	pong, err := Client.Ping(Ctx).Result()
	if err != nil {
		log.Fatalln("No Connection to Redis\n", err)
	}

	fmt.Println("Redis Connection is Established", pong)
}
