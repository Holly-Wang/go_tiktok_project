package service

import (
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/golang-jwt/jwt"
	"go_tiktok_project/common/dal/mysql"
	"regexp"
	"time"
)

// check whether username is legal
func IsUsernameLegal(username string) bool {

	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(username)

}

// check whether username exists
// true, nil --> user exists
// true, err --> error
// false, nil --> user doesn't exist
func IsUsernameExist(username string) (bool, error) {
	res, err := mysql.FindUserByName(username)
	if err != nil {
		// database error
		logs.Errorf("mysql error: ", err.Error())
		return true, err
	}
	return res, nil
}

// GenerateToken TODO: find right place for generate token
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
