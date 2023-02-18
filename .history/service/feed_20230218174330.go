package service

import (
	"context"
	"go_tiktok_project/common/authenticate"
	model "go_tiktok_project/common/dal/mysql"
	pb_feed "go_tiktok_project/idl/pb_feed"
	"time"
)

type Reponse struct {
	StatusCode    int64  `gorm:"status_code"`
	StatusMessage string `status`
}

type FeedResponse struct {
	Reponse
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

func GetFeedInfo(ctx context.Context, req *pb_feed.DouyinFeedRequest, userInfo *authenticate.UserInfo) (*pb_feed.DouyinFeedResponse, error) {
	var (
		resp            = new(pb_feed.DouyinFeedResponse)
		VideoReturnList = []*pb_feed.Video{}
	)
	var SuccessCode int32 = 1
	var FailCode int32 = 0
	var StatusMessage string = "1"
	model.InitDB()
	videos := []Video{}
	var Auther User
	var userName string
	var userId int64
	userName = userInfo.Username
	userId, err := model.FinduserNameById(userName)
	if err != nil {
		return &pb_feed.DouyinFeedResponse{
			StatusCode: &FailCode,
			StatusMsg:  &StatusMessage,
		}, err
	}
	video_sql, err_2 := model.FindVideoList()
	if err_2 != nil {
		return &pb_feed.DouyinFeedResponse{
			StatusCode: &FailCode,
			StatusMsg:  &StatusMessage,
		}, err_2
	}

	for _, v := range video_sql {
		video := &Video{
			Id: v.VideoID,
			//Author:         *v.Author,
			Play_url:       v.PlayUrl,
			Cover_url:      v.CoverUrl,
			Favorite_count: v.LikeCount,
			Comment_count:  v.CommentCount,
			//Is_favorite:    v.isLike,
			Title:    v.Title,
			Abstract: v.Abstract,
		}
		videos = append(videos, *video)
	}

	for i := 0; i < len(videos); i++ {
		Auther_ID := videos[i].Author.Id
		Video_ID := videos[i].Id
		var isLike bool
		var isFollow bool
		//	Video_ID := videos[i].VideoID
		isFollow, err := model.FindFollow(userId, Auther_ID)
		if err != nil {
			return &pb_feed.DouyinFeedResponse{
				StatusCode: &FailCode,
				StatusMsg:  &StatusMessage,
			}, err
		}
		isLike, err_2 := model.FindLike(userId, Video_ID)
		if err_2 != nil {
			return &pb_feed.DouyinFeedResponse{
				StatusCode: &FailCode,
				StatusMsg:  &StatusMessage,
			}, err_2
		}
		Auther, err_3 := model.FindUserInfoinUser(Auther_ID)
		if err_3 != nil {
			return &pb_feed.DouyinFeedResponse{
				StatusCode: &FailCode,
				StatusMsg:  &StatusMessage,
			}, err_3
		}
		var UserReturn pb_feed.User = pb_feed.User{
			Id:            &videos[i].Author.Id,
			Name:          &Auther.Username,
			FollowCount:   &Auther.Follow_cnt,
			FollowerCount: &Auther.Follower_cnt,
			IsFollow:      &isFollow,
		}
		print(*UserReturn.IsFollow)
		//user_id := auther.UserID
		var VideoResponse pb_feed.Video = pb_feed.Video{
			Id:            &videos[i].Id,
			Author:        &UserReturn,
			PlayUrl:       &videos[i].Play_url,
			CoverUrl:      &videos[i].Cover_url,
			FavoriteCount: &videos[i].Favorite_count,
			CommentCount:  &videos[i].Comment_count,
			Title:         &videos[i].Title,
			IsFavorite:    &isLike,
		}
		VideoReturnList = append(VideoReturnList, &VideoResponse)
		//fmt.Println(auther)
	}
	var Nexttime int64
	Nexttime = time.Now().Unix()
	resp.StatusCode = &SuccessCode
	resp.VideoList = VideoReturnList
	resp.NextTime = &Nexttime
	resp.StatusMsg = &StatusMessage

	return resp, nil
}
