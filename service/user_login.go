package service

import (
	"context"
	"go_tiktok_project/common/authenticate"
	"go_tiktok_project/common/dal/mysql"
	"go_tiktok_project/idl/biz/model/pb"
)

func UserLogin(ctx context.Context, req *pb.DouyinUserLoginRequest) (*pb.DouyinUserLoginResponse, error) {
	user, err := mysql.FindUserByNameAndPass(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	// TODO(liuyiyang): 暂时无 redis 环境
	// token, err := rediss.GetTokenByName(ctx, req.Username)
	// if err != nil {
	// 	return nil, err
	// }

	token, err := authenticate.GenToken(user.UserID, user.Username)
	if err != nil {
		return nil, err
	}

	resp := &pb.DouyinUserLoginResponse{
		UserId: user.UserID,
		Token:  token,
	}
	return resp, nil
}
