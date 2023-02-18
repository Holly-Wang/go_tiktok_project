package TokenCheck

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	SECRETKEY = "243223ffslsfsldfl412fdsfsdf"
)

type CustomClaims struct {
	UserName string
	PassWord string
	jwt.StandardClaims
}

func GenToken(userName string, passWord string) (string, error) {
	maxAge := 60 * 60 * 24
	claims := &CustomClaims{
		UserName: userName,
		PassWord: passWord,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(), // 过期时间，必须设置
			Issuer:    userName,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		fmt.Println(err)
	}
	return tokenString, nil
}
