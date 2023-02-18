package mysql

import (
	"time"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
)

// CreateLike 根据雪花主键插入视频的点赞表
func CreateLike(keyID, userID, videoID int64) error {
	like := Like{KeyID: keyID, OwnerID: userID, VideoID: videoID, LikeTime: time.Now()}
	if err := db.Create(&like).Error; err != nil {
		logs.Error("Like表格创建数据失败, err: %v", err)
		return err
	}
	return nil
}

func CreateUser(username, password string) (int64, error) {
	user := User{
		Username:     username,
		Password:     password,
		Follower_cnt: 0,
		Follow_cnt:   0,
		RegisterTime: time.Now(),
	}
	if err := db.Select("Username", "Password", "Follower_cnt", "Follow_cnt", "RegisterTime").Create(&user).Error; err != nil {
		logs.Errorf("insert to mysql err: %v", err)
		return 0, err
	}
	return user.UserID, nil
}

func CreateComment(keyId, videoID, userID int64, context string, likeCount int64, isLike bool) error {
	comment := Comment{CommentID: keyId, VideoID: videoID, UserID: userID,
		Context: context, LikeCount: likeCount, IsLike: isLike, CommentTime: time.Now()}
	if err := db.Create(&comment).Error; err != nil {
		logs.Error("Comment表格创建数据失败: %v", err)
		return err
	}
	return nil
}

// CreateVideo 根据雪花主键插入视频表
func CreateVideo(videoID, autherID int64, play_url, cover_url string, likecount, commentcount int64, title, abstract string) error {
	video := Video{VideoID: videoID, 
		AutherID: autherID,
		PlayUrl: play_url, 
		CoverUrl: cover_url,
		LikeCount: likecount, 
		CommentCount: commentcount, 
		Title: title, 
		Abstract: abstract,
	}
	// 通过数据的指针来创建
	result := db.Create(&video)
	if result.Error != nil {
		fmt.Println("Video表格创建数据失败: " + result.Error.Error())
		return result.Error
	}
	return nil
}
