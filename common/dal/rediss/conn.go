package rediss

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

var rdb *redis.Client

func getRedisInstance() *redis.Client {
	if rdb == nil {
		initRedis()
	}
	return rdb
}

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	// used for test
	//rdb = redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "",
	//	DB:       15,
	//})
	temp, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Println("redis connect failed + ", err.Error())
		return
	}
	log.Println("redis init success")

	_ = temp
}
