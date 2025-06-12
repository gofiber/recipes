package config

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func RedisInit() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})

	fmt.Println("Redis initialized")
}
