package rediss

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func GetRedis() *redis.Client {
	return rdb
}

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	temp, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Println("redis connetc failed + ", err.Error())
		return
	}
	log.Println("redis init success")

	_ = temp
}
