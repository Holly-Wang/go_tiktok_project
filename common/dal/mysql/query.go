package mysql

import "fmt"

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
