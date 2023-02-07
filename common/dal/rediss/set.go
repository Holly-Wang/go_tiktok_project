package rediss

import (
	"context"
	"errors"
	"log"
)

// key-value
// SET key value
func SetValue(key string, value interface{}) error {
	var ctx context.Context
	if key == "" {
		return errors.New("argument is null")
	}
	rdb := GetRedis()

	err := rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		return errors.New(err.Error())
	}
	log.Println("SET ", key, " ", value)
	return nil
}

// HSET key field value
// store token
// key: "toekn"
// field: username
// value: token
func HashSetToken(key, field string, value interface{}) error {
	var ctx context.Context
	if key == "" || field == "" {
		return errors.New("argument is null")
	}
	rdb := GetRedis()

	err := rdb.HSet(ctx, key, field, value).Err()
	if err != nil {
		return errors.New(err.Error())
	}
	log.Println("HSET ", key, " ", field, " ", value)
	return nil
}
