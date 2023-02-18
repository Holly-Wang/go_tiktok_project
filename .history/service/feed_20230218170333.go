package service

import (
	"context"
	"fmt"
	"go_tiktok_project/common/authenticate"
	model "go_tiktok_project/common/dal/mysql"
	pb_feed "go_tiktok_project/idl/pb_feed"
	"time"
)

type User struct {
	UserID       int64 `gorm:"primaryKey"`
	Username     string
	Password     string
	NickName     string
	Follower_cnt int64
	Follow_cnt   int64
	RegisterTime time.Time
}

type Video struct {
	VideoID      int64 `gorm:"primaryKey"`
	AutherID     int64
	PlayUrl      string
	CoverUrl     string
	LikeCount    int64
	CommentCount int64
	Title        string
	Abstract     string
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

func GetFeedInfo(ctx context.Context, req *pb_feed.DouyinFeedRequest, userInfo *authenticate.UserInfo) (*pb_feed.DouyinFeedResponse, error) {
	var (
		resp            = new(pb_feed.DouyinFeedResponse)
		VideoReturnList = []*pb_feed.Video{}
	)
	model.InitDB()
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
		var UserReturn pb_feed.User = pb_feed.User{
			Id:            &videos[i].AutherID,
			Name:          &auther.UserName,
			FollowCount:   &auther.FollowCount,
			FollowerCount: &auther.FollowerCount,
			IsFollow:      &isFollow,
		}
		print(*UserReturn.IsFollow)
		//user_id := auther.UserID
		var VideoResponse pb_feed.Video = pb_feed.Video{
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
