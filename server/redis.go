package server

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func setupRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
