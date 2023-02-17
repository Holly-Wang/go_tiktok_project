package mysql

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
)

// FindIDinLike 失败时主键返回0和错误信息
func FindIDinLike(userID, videoID uint64) (int64, error) {
	var like Like
	if err := db.Where("owner_id = ? AND video_id = ?", userID, videoID).First(&like).Error; err.Error != nil {
		logs.Error("查询like表主键出错, err: %v", err.Error())
		return 0, err
	}
	return like.KeyID, nil
}

// CheckUserExist check whether user exist
func CheckUserExist(username string) (bool, error) {
	var count int64
	if err := db.Model(&User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func FindUserByNameAndPass(username, password string) (User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, errors.New("user doesn't exist")
	}
	if err != nil {
		logs.Errorf("mysql error during selecting: ", err.Error())
		return user, err
	}
	if user.Password != password {
		return user, errors.New("wrong password")
	}
	return user, nil
}

func FindUserById(userid uint64) (User, error) {
	var user User
	err := db.Where("user_id = ?", userid).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return user, errors.New("user doesn't exist")
	}
	if err != nil {
		logs.Errorf("mysql error during selecting: ", err.Error())
		return user, err
	}
	return user, nil
}

func FindComment(videoID int64) ([]Comment, error) {
	var comments []Comment
	// select * from comments where video_id = ? order by comment_time desc
	res := db.Where("video_id = ? ", videoID).Order("comment_time desc").Find(&comments)
	if res.Error != nil {
		logs.Error("查询comment表主键出错, err: ", res.Error)
		return nil, res.Error
	}
	var commentss []Comment
	for i := 0; i < int(res.RowsAffected); i++ {
		commentss = append(commentss, comments[i])
	}
	return commentss, nil
}

// FindLikeList 查询登录用户喜欢的视频列表
func FindLikeList(userID int64) (*[]int64, error) {
	var likes []Like
	res := db.Where("owner_id = ?", userID).Find(&likes)
	if res.Error != nil {
		logs.Error("无法获取用户喜爱列表, err: %V", res.Error.Error())
		return nil, res.Error
	}
	var videoIDs []int64
	for i := 0; i < int(res.RowsAffected); i++ {
		videoIDs = append(videoIDs, likes[i].VideoID)
	}
	return &videoIDs, nil
}

// FindLikeOfVideo 查询视频点赞数
func FindLikeOfVideo(videoID int64) (int64, error) {
	var video Video
	err := db.Where("video_id = ?", videoID).First(&video).Error
	if err != nil {
		fmt.Println("查询点赞出错, error: " + err.Error())
		return -1, err
	}
	return video.LikeCount, nil
}

// CheckFollow 校验source是否关注target
func CheckFollow(sourceID, targetID int64) (bool, error) {
	var count int64
	if err := db.Model(&Follow{}).Where("watcher_id = ? and watched_id = ?", sourceID, targetID).
		Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
