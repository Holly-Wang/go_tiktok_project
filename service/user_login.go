package service

import (
	"context"
	"go_tiktok_project/common/dal/mysql"
	"go_tiktok_project/common/dal/rediss"
	"go_tiktok_project/idl/pb"
)

func UserLogin(ctx context.Context, req *pb.DouyinUserLoginRequest) (*pb.DouyinUserLoginResponse, error) {
	user, err := mysql.FindUserByNameAndPass(*req.Username, *req.Password)
	if err != nil {
		return nil, err
	}

	token, err := rediss.GetTokenByName(ctx, *req.Username)
	if err != nil {
		return nil, err
	}

	resp := &pb.DouyinUserLoginResponse{
		StatusCode: new(int32),
		StatusMsg:  new(string),
		UserId:     &user.UserID,
		Token:      &token,
	}
	return resp, nil
}
