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

func FindUserById(userID uint64) (User, error) {
	var user User
	err := db.Where("user_id = ?", userID).First(&user).Error
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
func FindLikeList(userID int64) ([]int64, error) {
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
	return videoIDs, nil
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

// 根据查询userID发布的视频列表
func FindVideoListinVideo(userID int64) ([]*Video, error) {
	var Video_list []*Video
	err := db.Where("auther_id = ?", userID).Find(&Video_list)

	if err.Error != nil {
		fmt.Println("查询Videl表主键出错, error: " + err.Error.Error())
		return nil, err.Error
	}

	return Video_list, nil
}

// 查询userID对应的user信息
func FindUserInfoinUser(userID int64) (User, error) {
	var user User
	err := db.Where("user_id = ?", userID).Find(&user)

	if err.Error != nil {
		fmt.Println("查询User表主键出错, error: " + err.Error.Error())
		return user, err.Error
	}

	return user, nil
}

// 查询user给视频的点赞次数
func FindCountinLikes(userID int64, videoID int64) (int64, error) {
	var total int64 = 0

	err := db.Where("owner_id = ? AND video_id = ?", userID, videoID).Model(Like{}).Count(&total)
	if err.Error != nil {
		fmt.Println("查询like表主键出错, error: " + err.Error.Error())
		return -1, err.Error
	}

	return total, nil
}

// 查询视频表中最大的vedio_id
func FindMaxIdinVideos() (int64, error) {
	var video Video

	err := db.Order("video_id desc").First(&video)
	if err.Error != nil {
		fmt.Println("查询Video表最大id出错, error: " + err.Error.Error())
		return -1, err.Error
	}

	return video.VideoID, nil
}

// 查询wachter是否关注watched
func FindCountinFollows(watcher_id int64, watched_id int64) (int64, error) {
	var total int64 = 0

	err := db.Where("watcher_id = ? AND watched_id = ?", watcher_id, watched_id).Model(Follow{}).Count(&total)
	if err.Error != nil {
		fmt.Println("查询follows表主键出错, error: " + err.Error.Error())
		return -1, err.Error
	}

	return total, nil
}

// 返回视频列表
func FindVideoList() ([]*Video, error) {
	var Video_list []*Video
	err := db.Raw("select * from videos order by video_id limit 30").Scan(&Video_list)
	if err.Error != nil {
		fmt.Println("获取视频列表错误, error: " + err.Error.Error())
		return Video_list, err.Error
	}

	return Video_list, nil
}

// 根据username找id
func FinduserNameById(userName string) (int64, error) {
	var userId int64
	err := db.Raw("SELECT user_id FROM `users` WHERE username=?", userName).Scan(&userId)
	if err.Error != nil {
		fmt.Println("获取用户ID出错, error: " + err.Error.Error())
		return -1, err.Error
	}

	return userId, nil
}

// 查询follow关系
func FindFollow(user_1 int64, user_2 int64) (bool, error) {
	var isFollow bool
	var countFollow int64
	sql_follow := "SELECT count(key_id) FROM `follows` WHERE watcher_id=? and watched_id=?"
	err := db.Raw(sql_follow, user_1, user_2).Scan(&countFollow)
	if err.Error != nil {
		fmt.Println("获取关注关系出错, error: " + err.Error.Error())
		return false, err.Error
	}
	if countFollow > 0 {
		isFollow = true
	} else {
		isFollow = false
	}
	return isFollow, nil
}

// 查询喜欢视频关系
func FindLike(userId int64, videoId int64) (bool, error) {
	var isLike bool
	var countLike int64
	sql_like := "SELECT count(key_id) FROM `likes` WHERE owner_id=? and video_id=?"
	err := db.Raw(sql_like, userId, videoId).Scan(&countLike)
	if err.Error != nil {
		fmt.Println("获取喜欢关系出错, error: " + err.Error.Error())
		return false, err.Error
	}
	if countLike > 0 {
		isLike = true
	} else {
		isLike = false
	}
	return isLike, nil
}
