package TokenCheck

import (
	"context"
	"fmt"
	"go_tiktok_project/common/authenticate"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CheckToken(tokenString string) (string, error) {
	token, err_1 := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETKEY), nil
	})
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		sql_check := "select count(user_id) from users where username=? and password=?"
		var userName string
		var passWord string
		var isExist int64
		userName = claims.UserName
		passWord = claims.PassWord
		url := "MiniTikTok:root@tcp(49.232.155.203:3306)/minitiktok?charset=utf8&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
		if err != nil {
			fmt.Println("连接失败")
			return "false", fmt.Errorf("Error in database!")
		} else {
			fmt.Println("连接成功")
		}
		db.Raw(sql_check, userName, passWord).Scan(&isExist)
		if isExist == 1 {
			return "true", nil
		} else {
			return "false", fmt.Errorf("Unexisted username or password!")
		}
	} else {
		return "false", err_1
	}
}

func TokenCheck() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// TODO(...): 添加登录验证并解析用户信息
		// 验证失败： c.Abort()
		// 验证成功: c.Next(ctx)
		//          c.Set(登录用户信息)

		// Mock
		tokenString := c.Query("token")
		logs.Info(tokenString)
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(SECRETKEY), nil
		})
		if err != nil {
			logs.Info("111")
			c.String(400, "Token Error!!!")
			c.Abort()
		}
		var userName string
		var passWord string
		var isExist int64
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			sql_check := "select count(user_id) from users where username=? and password=?"
			userName = claims.UserName
			passWord = claims.PassWord
			url := "MiniTikTok:root@tcp(49.232.155.203:3306)/minitiktok?charset=utf8&parseTime=True&loc=Local"
			db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
			if err != nil {
				fmt.Println("连接失败")
				c.String(400, "Connection Error!")
				c.Abort()
			} else {
				fmt.Println("连接成功")
			}
			db.Raw(sql_check, userName, passWord).Scan(&isExist)
			if isExist != 1 {
				c.String(400, "Unexisted username or password!")
				c.Abort()
			}
		} else {
			c.String(400, "Unexisted username or password!")
			c.Abort()
		}

		userName_get := c.Query("user_name")
		if userName_get != userName {
			c.String(400, "Username inconsistent!")
			c.Abort()

		}
		//		_, err := CheckToken(Token)
		//		if err != nil {
		//			c.String(400, "Token Error!")
		//			c.Abort()
		//		}
		c.Set(authenticate.ReqUserInfoKey, &authenticate.UserInfo{
			Username: userName,
		})
		c.Next(ctx)
	}
}
