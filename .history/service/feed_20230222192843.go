package service

import (
	"context"
	"go_tiktok_project/common/authenticate"
	model "go_tiktok_project/common/dal/mysql"
	pb "go_tiktok_project/idl/biz/model/pb"
	"time"
)

type Reponse struct {
	StatusCode    int64  `gorm:"status_code"`
	StatusMessage string `json:"status_message"`
}

type FeedResponse struct {
	Reponse
	VideoList []pb.Video `json:"video_list,omitempty"`
	NextTime  int64      `json:"next_time,omitempty"`
}

func GetFeedInfo(ctx context.Context, req *pb.DouyinFeedRequest, userInfo *authenticate.UserInfo, isLogin bool) (*pb.DouyinFeedResponse, error) {
	var (
		resp            = new(pb.DouyinFeedResponse)
		VideoReturnList = []*pb.Video{}
		videos          = []*pb.Video{}
	)
	var FailCode int32 = 1
	var StatusMessage string = "1"
	var userId int64
	if isLogin == true {
		userId = userInfo.UserID
	}
	video_sql, err := model.FindVideoList()
	if err != nil {
		return &pb.DouyinFeedResponse{
			StatusCode: FailCode,
			StatusMsg:  StatusMessage,
		}, err
	}

	for _, v := range video_sql {
		video := &pb.Video{
			Id: v.VideoID,
			Author: &pb.User{
				Id: v.AutherID,
			},
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.LikeCount,
			CommentCount:  v.CommentCount,
			//IsFavorite:    v.isLike,
			Title: v.Title,
		}
		videos = append(videos, video)
	}

	for i := 0; i < len(videos); i++ {
		Auther_ID := videos[i].Author.Id
		Video_ID := videos[i].Id
		//获取作者id和视频id
		var isFollow bool//表示用户(若登录)是否关注当前视频作者
		var isLike bool//表示用户(若登录)是否喜欢当前视频
		isFollow = false
		isLike = false
		if isLogin == true {//只在用户登录的情况下查询关注(作者)和喜欢(视频)信息
			isFollow, err = model.FindFollow(userId, Auther_ID)//查询关注信息
			if err != nil {
				return &pb.DouyinFeedResponse{
					StatusCode: FailCode,
					StatusMsg:  StatusMessage,
				}, err
			}
			isLike, err = model.FindLike(userId, Video_ID)//查询喜欢信息
			if err != nil {
				return &pb.DouyinFeedResponse{
					StatusCode: FailCode,
					StatusMsg:  StatusMessage,
				}, err
			}
		}
		Auther, err := model.FindUserInfoinUser(Auther_ID)//查询作者信息
		if err != nil {
			return &pb.DouyinFeedResponse{
				StatusCode: FailCode,
				StatusMsg:  StatusMessage,
			}, err
		}
		var UserReturn pb.User = pb.User{//构造作者信息
			Id:            videos[i].Author.Id,
			Name:          Auther.Username,
			FollowCount:   Auther.Follow_cnt,
			FollowerCount: Auther.Follower_cnt,
			IsFollow:      isFollow,
		}
		VideoResponse := pb.Video{//构造视频信息
			Id:            videos[i].Id,
			Author:        &UserReturn,
			PlayUrl:       videos[i].PlayUrl,
			CoverUrl:      videos[i].CoverUrl,
			FavoriteCount: videos[i].FavoriteCount,
			CommentCount:  videos[i].CommentCount,
			Title:         videos[i].Title,
			IsFavorite:    isLike,
		}
		VideoReturnList = append(VideoReturnList, &VideoResponse)
	}
	resp.VideoList = VideoReturnList
	resp.NextTime = time.Now().Unix()
	return resp, nil
}
