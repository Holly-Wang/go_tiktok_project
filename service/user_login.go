package service

import (
	"context"
	"errors"
	"go_tiktok_project/common/authenticate"
	"go_tiktok_project/common/dal/mysql"
	"go_tiktok_project/idl/biz/model/pb"

	"golang.org/x/crypto/bcrypt"
)

func comparePwd(pwd1 string, pwd2 string) bool {
	// Returns true on success, pwd1 is for the database.
	err := bcrypt.CompareHashAndPassword([]byte(pwd1), []byte(pwd2))
	if err != nil {
		return false
	} else {
		return true
	}
}

func UserLogin(ctx context.Context, req *pb.DouyinUserLoginRequest) (*pb.DouyinUserLoginResponse, error) {
	user, err := mysql.FindUserByNameAndPass(req.Username)
	if err != nil {
		return nil, err
	}

	if !comparePwd(user.Password, req.Password) {
		return nil, errors.New("wrong password")
	}

	// TODO(liuyiyang): 暂时无 redis 环境
	// token, err := rediss.GetTokenByName(ctx, req.Username)
	// if err != nil {
	// 	return nil, err
	// }

	token, err := authenticate.GenToken(user.UserID, user.Username)

	resp := &pb.DouyinUserLoginResponse{
		UserId: user.UserID,
		Token:  token,
	}
	return resp, nil
}
