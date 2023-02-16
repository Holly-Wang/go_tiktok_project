package handler

import (
	"context"
	"go_tiktok_project/common/authenticate"
	pb "go_tiktok_project/idl/pb"
	"go_tiktok_project/service"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/cloudwego/hertz/pkg/app"
)

func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()
	logs.Info("req path: %s", path)

	req := new(pb.DouyinUserRequest)
	if err := c.BindAndValidate(&req); err != nil {
		c.String(400, err.Error())
		return
	}

	userInfo, err := authenticate.GetAuthUserInfo(c)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp, err := service.GetUserInfo(ctx, req, userInfo)
	if err != nil {
		logs.Errorf("service err: %v", err)
		c.String(400, err.Error())
		return
	}

	c.JSON(200, resp)
}
