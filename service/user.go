package service

import (
	"context"
	"go_tiktok_project/common/authenticate"
	"go_tiktok_project/idl/pb"
)

func GetUserInfo(ctx context.Context, req *pb.DouyinUserRequest, userInfo *authenticate.UserInfo) (*pb.DouyinUserResponse, error) {
	var (
		resp = new(pb.DouyinUserResponse)
	)
	return resp, nil
}
