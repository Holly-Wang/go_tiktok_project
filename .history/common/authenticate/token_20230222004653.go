package authenticate

import (
	"fmt"
	"time"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/golang-jwt/jwt"
)

const (
	SECRETKEY = "243223ffslsfsldfl412fdsfsdf"
)

type CustomClaims struct {
	UserID   int64
	UserName string
	jwt.StandardClaims
}

func GenToken(userID int64, userName string) (string, error) {
	maxAge := 60 * 60 * 24
	claims := &CustomClaims{
		UserID:   userID,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(),
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

func CheckToken(tokenString string) (*UserInfo, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	if err != nil {
		logs.Error("token err: %v", err)
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok || !token.Valid {
		logs.Error("token check err: %v", err)
		return nil, err
	}
	return &UserInfo{
		UserID:   claims.UserID,
		Username: claims.UserName,
	}, nil
}
