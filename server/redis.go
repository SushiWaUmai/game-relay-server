package server

import (
	"log"
	"os"

	goredis "github.com/go-redis/redis"
	"github.com/nitishm/go-rejson"
)

var RedisClient *goredis.Client
var RedisJsonHandler *rejson.Handler

func init() {
	addr := "localhost:6379"

	RedisClient = goredis.NewClient(&goredis.Options{
		Addr: addr,
	})

	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	RedisJsonHandler := rejson.NewReJSONHandler()
	RedisJsonHandler.SetGoRedisClient(RedisClient)
}

func CheckRedisError(res interface{}, err error) {
	if err != nil {
		log.Fatalf("Failed to JSONSet")
		return
	}

	if res.(string) == "OK" {
		log.Printf("Success: %s\n", res)
	} else {
		log.Println("Failed to Set: ")
	}
}
