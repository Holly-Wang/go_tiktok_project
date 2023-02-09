package mysql

import (
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/redis/go-redis/v9"
)

// FindIDinLike 失败时主键返回0和错误信息
func FindIDinLike(userID, videoID uint64) (uint64, error) {
	var like Like
	err := db.Where("owner_id = ? AND video_id = ?", userID, videoID).First(&like)
	if err.Error != nil {
		fmt.Println("查询like表主键出错, error: " + err.Error.Error())
		return 0, err.Error
	}
	return like.KeyID, nil
}

// FindUserByName check whether username exist
// false, nil --> user isn't found
// false, err --> database error
// true, nil --> user is found
func FindUserByName(username string) (bool, error) {
	var user User
	err := db.Where("username = ?", username).First(&user)
	if err.Error == redis.Nil {
		// user isn't found
		return false, nil
	}
	// other error
	if err != nil {
		return false, err.Error
	}
	// user is found
	return true, nil
}

func FindUserByNameAndPass(username, password string) (User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user)
	if err.Error == redis.Nil {
		return user, errors.New("user doesn't exist")
	}
	if err != nil {
		logs.Errorf("mysql error during selecting: ", err.Error.Error())
		return user, err.Error
	}
	if user.Password != password {
		return user, errors.New("wrong password")
	}
	return user, nil
}
