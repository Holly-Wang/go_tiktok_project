package common

import (
	"fmt"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	ID       uint64
	UserName string
	jwt.StandardClaims
}

var jwtkey string = "123456"

func ParseToken(token string) (*Claims, error) {
	// ParseWithClaims 解析token
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 使用签名解析用户传入的token,获取载荷部分数据
		return []byte(jwtkey), nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		//Valid用于校验鉴权声明。解析出载荷部分
		if c, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return c, nil
		}
	}
	return nil, err
}

func Token2UserID(token string) (uint64, error) {
	// 非法token返回0 + error
	var claimp = new(Claims)
	var err error

	claimp, err = ParseToken(token)
	if err != nil {
		fmt.Println("token is illegal: " + err.Error())
		return 0, err
	}
	return claimp.ID, nil
}
