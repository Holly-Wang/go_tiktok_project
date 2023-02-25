package service

import (
	"context"
	"go_tiktok_project/common/authenticate"
	"go_tiktok_project/common/dal/mysql"
	"go_tiktok_project/idl/biz/model/pb"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
)

func GetUserInfo(ctx context.Context, req *pb.DouyinUserRequest, userInfo *authenticate.UserInfo) (*pb.DouyinUserResponse, error) {
	userRecord, err := mysql.FindUserById(uint64(userInfo.UserID))
	if err != nil {
		logs.Errorf("[SQL Error] get user record err: %v", err)
		return nil, err
	}

	//isFollow, err := mysql.CheckFollow(userInfo.UserID, req.UserId)
	//if err != nil {
	//	logs.Errorf("[SQL Error] check follow err: %v", err)
	//	return nil, err
	//}

	resp := &pb.DouyinUserResponse{
		User: &pb.User{
			Id:            req.UserId,
			Name:          userRecord.NickName,
			FollowCount:   userRecord.Follow_cnt,
			FollowerCount: userRecord.Follower_cnt,
			IsFollow:      false,
		},
	}
	return resp, nil
}
