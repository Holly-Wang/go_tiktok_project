package service

import (
	"context"
	"fmt"
	"go_tiktok_project/common"
	"go_tiktok_project/common/authenticate"
	"go_tiktok_project/common/dal/mysql"
	"go_tiktok_project/idl/biz/model/pb"
	"regexp"

	"golang.org/x/crypto/bcrypt"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
)

const (
	pattern = `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
)

var (
	reg = regexp.MustCompile(pattern)
)

func generateCipher(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return hash, err
}

func UserRegister(ctx context.Context, req *pb.DouyinUserRegisterRequest) (*pb.DouyinUserRegisterResponse, error) {
	if err := checkRegisterUser(req.Username); err != nil {
		return nil, err
	}

	cipher, err := generateCipher(req.Password)
	if err != nil {
		return nil, err
	}

	userID, err := mysql.CreateUser(req.Username, string(cipher))
	if err != nil {
		return nil, err
	}

	token, err := authenticate.GenToken(userID, req.Username)
	if err != nil {
		return nil, err
	}

	//if err := rediss.SetToken(ctx, req.Username, token); err != nil {
	//	return nil, err
	//}

	resp := &pb.DouyinUserRegisterResponse{
		StatusMsg:  common.RegisterSueecssMsg,
		UserId:     userID,
		Token:      token,
		StatusCode: common.RegisterSucces,
	}
	return resp, nil
}

// checkRegisterUser check register user (username legal and whether has exist)
func checkRegisterUser(username string) error {
	logs.Infof("check register user")
	// check username legal
	if !reg.MatchString(username) {
		logs.Error("username illegal")
		return fmt.Errorf("username is illegal")
	}
	// check user has exist
	if exist, err := mysql.CheckUserExist(username); err != nil || exist {
		logs.Error("user has exist")
		if err == nil {
			err = fmt.Errorf("user has exist")
		}
		return err
	}
	return nil
}
