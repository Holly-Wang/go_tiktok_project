package mysql

import (
	"fmt"

	"gorm.io/gorm"
)

// UpdateLikeCount 视频增加点赞后必须马上调用
func UpdateLikeCount(videoID uint64) error { 
	result := db.Model(&Video{}).Where("video_id=?", videoID).Update("like_count", gorm.Expr("like_count+?", 1))
	if result.Error != nil {
		fmt.Println("Video表格LikeCount列更新出错" + result.Error.Error())
		return result.Error
	}
	return nil
}
