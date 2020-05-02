package redis

import (
	"log"

	"gin-app-start/app/config"

	"github.com/go-redis/redis"
)

var Redis *redis.Client

func Init() {
	config := config.Conf.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password, // no password set
		DB:       config.Db,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Println("redis connection error: ", err)
		panic(err)
	}
	log.Panicln("redis connect ping response:", pong)
	Redis = client
}
