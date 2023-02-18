package handler

import (
	"context"
	"go_tiktok_project/common/authenticate"
	pb_feed "go_tiktok_project/idl/pb_feed"
	"go_tiktok_project/service"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/cloudwego/hertz/pkg/app"
)

func GetFeedInfo(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()
	logs.Info("req path: %s", string(path))

	req := new(pb_feed.DouyinFeedRequest)
	if err := c.BindAndValidate(&req); err != nil {
		c.String(400, err.Error())
		return
	}
	var userInfo *authenticate.UserInfo
	var userInfo_get, err_bool = c.Get(authenticate.ReqUserInfoKey)
	if err_bool != true {
		c.String(400, "No UserInfo!")
		return
	}
	userInfo = userInfo_get.(*authenticate.UserInfo)

	resp, err := service.GetFeedInfo(ctx, req, userInfo)
	if err != nil {
		logs.Errorf("service err: %v", err)
		c.String(400, err.Error())
		return
	}

	c.JSON(200, resp)
}
