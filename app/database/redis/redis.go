package redis

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/onsi/ginkgo/config"
)

var Redis *redis.Client

func Init() {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password, // no password set
		DB:       config.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Println("redis connection error: ", err)
		panic(err)
	}
	log.Panicln("redis connect ping response:", pong)
	Redis = client
}
