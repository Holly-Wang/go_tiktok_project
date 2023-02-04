package middlewares

import (
	"context"
	"go_tiktok_project/common/authenticate"

	"github.com/cloudwego/hertz/pkg/app"
)

func AuthN() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// TODO(...): 添加登录验证并解析用户信息
		// 验证失败： c.Abort()
		// 验证成功: c.Next(ctx)
		//          c.Set(登录用户信息)

		// Mock
		c.Set(authenticate.ReqUserInfoKey, &authenticate.UserInfo{
			Username: "test_user",
		})
	}
}
