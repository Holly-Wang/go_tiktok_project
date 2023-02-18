package service

import (
	"context"
	"fmt"
	"go_tiktok_project/common/dal/mysql"
	"go_tiktok_project/common/dal/rediss"
	"go_tiktok_project/idl/pb"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
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
	if err := checkRegisterUser(*req.Username); err != nil {
		return nil, err
	}

	cipher, err := generateCipher(*req.Password)
	if err != nil {
		return nil, err
	}

	userID, err := mysql.CreateUser(*req.Username, string(cipher))
	if err != nil {
		return nil, err
	}

	token, err := GenerateToken(uint64(userID), *req.Username)
	if err != nil {
		return nil, err
	}

	if err := rediss.SetToken(ctx, *req.Username, token); err != nil {
		return nil, err
	}

	resp := &pb.DouyinUserRegisterResponse{
		StatusCode: new(int32),
		StatusMsg:  new(string),
		UserId:     &userID,
		Token:      &token,
	}
	return resp, nil
}

// checkRegisterUser check register user (username legal and whether has exist)
func checkRegisterUser(username string) error {
	// check username legal
	if !reg.MatchString(username) {
		return fmt.Errorf("username is illegal")
	}
	// check user has exist
	if exist, err := mysql.CheckUserExist(username); err != nil || exist {
		if err == nil {
			err = fmt.Errorf("user has exist")
		}
		return err
	}
	return nil
}

func GenerateToken(userID uint64, username string) (string, error) {
	type Claims struct {
		ID                 uint64
		UserName           string
		jwt.StandardClaims // jwt中标准格式,主要是设置token的过期时间
	}

	var jwtkey = "123456"

	nowTime := time.Now()
	expireTime := nowTime.Add(300 * time.Second)
	issuer := "what2eat"
	claims := Claims{
		ID:       userID,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), // 转成纳秒
			Issuer:    issuer,
		},
	}
	// 根据签名生成token，NewWithClaims(加密方式,claims) ==> 头部，载荷，签证
	toke, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(jwtkey))
	return toke, err
}
