package redis

import (
	"gin-app-start/app/common"
	"gin-app-start/app/config"

	"github.com/go-redis/redis"
)

var logger = common.Logger
var Redis *redis.Client

func Init() {
	config := config.Conf.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password, // no password set
		DB:       config.Db,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		logger.Error("redis connection error: ", err)
		panic(err)
	}
	logger.Info("redis connect ping response:", pong)
	Redis = client
}
