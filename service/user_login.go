package service

import (
	"context"
	"errors"
	"go_tiktok_project/common/dal/mysql"
	"go_tiktok_project/common/dal/rediss"
	"go_tiktok_project/idl/pb"
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
	user, err := mysql.FindUserByNameAndPass(*req.Username)
	if err != nil {
		return nil, err
	}

	if !comparePwd(user.Password, *req.Password) {
		return nil, errors.New("wrong password")
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
