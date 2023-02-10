package mysql

import (
	"fmt"
	"time"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
)

// CreateLike 根据雪花主键插入视频的点赞表
func CreateLike(keyID, userID, videoID int64) error {
	like := Like{KeyID: keyID, OwnerID: userID, VideoID: videoID, LikeTime: time.Now()}
	// 通过数据的指针来创建
	result := db.Create(&like)
	if result.Error != nil {
		fmt.Println("Like表格创建数据失败: " + result.Error.Error())
		return result.Error
	}
	return nil
}

func CreateUser(username, password string) (int64, error) {
	var user User
	user = User{
		Username:     username,
		Password:     password,
		Follower_cnt: 0,
		Follow_cnt:   0,
		RegisterTime: time.Now(),
	}
	result := db.Create(&user)
	if result.Error != nil {
		logs.Errorf("insert to mysql error: ", result.Error.Error())
		//fmt.Println("User insert failed: " + result.Error.Error())
		return 0, result.Error
	}
	return user.UserID, nil
}
