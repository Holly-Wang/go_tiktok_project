package middlewares

import (
	"context"
	"go_tiktok_project/common/authenticate"

	"github.com/cloudwego/hertz/pkg/app"
)

func AuthN() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		token := c.Query("token")
		userInfo, err := authenticate.CheckToken(token)
		if err != nil {
			c.Abort()
			return
		}
		c.Set(authenticate.ReqUserInfoKey, userInfo)
		c.Next(ctx)
	}
}
