package service

import (
	"context"
	"fmt"
	"go_tiktok_project/common/authenticate"
	pb "go_tiktok_project/idl/pb"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Auther struct {
	UserID        int64  `gorm:"column:user_id"`
	UserName      string `gorm:"column:username"`
	FollowerCount int64  `gorm:"column:follower_cnt"`
	FollowCount   int64  `gorm:"column:follow_cnt"`
}

type Video_ struct {
	VideoID      int64  `gorm:"column:video_id"`
	AutherID     int64  `gorm:"column:auther_id"`
	PlayUrl      string `gorm:"column:play_url"`
	CoverUrl     string `gorm:"column:cover_url"`
	LikeCount    int64  `gorm:"column:like_count"`
	CommentCount int64  `gorm:"column:comment_count"`
	Title        string `gorm:"column:title"`
	Auther       `struct Auther`
}

type Reponse struct {
	StatusCode    int64  `gorm:"status_code"`
	StatusMessage string `status`
}

type FeedResponse struct {
	Reponse
	VideoList []Video_ `json:"video_list,omitempty"`
	NextTime  int64    `json:"next_time,omitempty"`
}

func GetFeedInfo(ctx context.Context, req *pb.DouyinFeedRequest, userInfo *authenticate.UserInfo) (*pb.DouyinFeedResponse, error) {
	var (
		resp            = new(pb.DouyinFeedResponse)
		VideoReturnList = []*pb.Video{}
	)

	//定义连接信息，用户名、密码、通信协议等信息
	url := "MiniTikTok:root@tcp(49.232.155.203:3306)/minitiktok?charset=utf8&parseTime=True&loc=Local"
	//连接数据库,连接数据库时，可以加上一些高级配置,就是gorm.Config中的参数
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		fmt.Println("连接失败")
		return resp, err
	} else {
		fmt.Println("连接成功")
	}
	sql_video := "select * from videos order by video_id limit 3"
	sql_auther := "select * from users where user_id =?"
	sql_follow := "SELECT count(key_id) FROM `follows` WHERE watcher_id=? and watched_id=?"
	sql_like := "SELECT count(key_id) FROM `likes` WHERE owner_id=? and video_id=?"
	sql_getid := "SELECT user_id FROM `users` WHERE username=?"
	videos := []Video_{}
	var auther Auther
	var countLike int64
	var countFollow int64
	var userName string
	var userId int64
	userName = userInfo.Username
	db.Raw(sql_getid, userName).Scan(&userId)
	db.Raw(sql_video).Scan(&videos)
	for i := 0; i < len(videos); i++ {
		Auther_ID := videos[i].AutherID
		Video_ID := videos[i].VideoID
		var isLike bool
		var isFollow bool
		//	Video_ID := videos[i].VideoID
		db.Raw(sql_follow, userId, Auther_ID).Scan(&countFollow)
		db.Raw(sql_like, userId, Video_ID).Scan(&countLike)
		db.Raw(sql_auther, Auther_ID).Scan(&auther)
		if countFollow > 0 {
			isFollow = true
		} else {
			isFollow = false
		}

		if countLike > 0 {
			isLike = true
		} else {
			isLike = false
		}
		var UserReturn pb.User = pb.User{
			Id:            &videos[i].AutherID,
			Name:          &auther.UserName,
			FollowCount:   &auther.FollowCount,
			FollowerCount: &auther.FollowerCount,
			IsFollow:      &isFollow,
		}
		print(*UserReturn.IsFollow)
		//user_id := auther.UserID
		var VideoResponse pb.Video = pb.Video{
			Id:            &videos[i].VideoID,
			Author:        &UserReturn,
			PlayUrl:       &videos[i].PlayUrl,
			CoverUrl:      &videos[i].CoverUrl,
			FavoriteCount: &videos[i].LikeCount,
			CommentCount:  &videos[i].CommentCount,
			Title:         &videos[i].Title,
			IsFavorite:    &isLike,
		}
		VideoReturnList = append(VideoReturnList, &VideoResponse)
		fmt.Println(auther)
	}
	var a int32
	a = 0
	var Nexttime int64
	var StatusMessage string
	StatusMessage = "1"
	Nexttime = time.Now().Unix()
	resp.StatusCode = &a
	resp.VideoList = VideoReturnList
	resp.NextTime = &Nexttime
	resp.StatusMsg = &StatusMessage

	return resp, nil
}
