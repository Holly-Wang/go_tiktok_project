package mysql

import (
	"fmt"
	"time"
)

func CreateLike(keyID uint64, userID uint64, videoID uint64) error { //根据雪花主键插入视频的点赞表
	like := Like{KeyID: keyID, OwnerID: userID, VideoID: videoID, LikeTime: time.Now()}
	result := db.Create(&like) // 通过数据的指针来创建
	if result.Error != nil {
		fmt.Println("Like表格创建数据失败: " + result.Error.Error())
		return result.Error
	}
	return nil
}
