package rediss

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"

	"github.com/redis/go-redis/v9"
)

// GetTokenByName get user's token from redis
func GetTokenByName(ctx context.Context, username string) (string, error) {
	key := "token"
	res, err := rdb.HGet(ctx, key, username).Result()
	if err == redis.Nil {
		return "", errors.New("Error occurs: key is not found.")
	}
	if err != nil {
		logs.Errorf("redis error: ", err.Error())
		return "", errors.New(err.Error())
	}
	return res, nil
}
