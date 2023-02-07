package rediss

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

// key-value
// GET key
func GetTVaule(key string) (string, error) {
	var ctx context.Context
	if key == "" {
		return "", errors.New("Error occurs: key is empty.")
	}
	rdb := GetRedis()
	res, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", errors.New("Error occurs: key is not found.")
	}
	if err != nil {
		return "", errors.New(err.Error())
	}
	return res, nil
}

// HGET key field
// key: "toekn"
// field: username
// value: token
func HashGetToken(key, field string) (string, error) {
	var ctx context.Context
	rdb := GetRedis()
	res, err := rdb.HGet(ctx, key, field).Result()
	if err == redis.Nil {
		return "", errors.New("Error occurs: key is not found.")
	}
	if err != nil {
		return "", errors.New(err.Error())
	}
	return res, nil
}

// HGETALL key
// use to get all tokens
// key: "toekn"
// field: username
// value: token
func HashGetAllTokens(key string) (map[string]string, error) {
	var ctx context.Context
	rdb := GetRedis()
	res, err := rdb.HGetAll(ctx, key).Result()
	if err == redis.Nil {
		return nil, errors.New("Error occurs: key is not found.")
	}
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return res, nil
}
