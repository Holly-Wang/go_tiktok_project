package handler

import (
	"context"
	"go_tiktok_project/common/authenticate"
	pb "go_tiktok_project/idl/biz/model/pb"
	"go_tiktok_project/service"
	"net/http"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/cloudwego/hertz/pkg/app"
)

func GetFeedInfo(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()
	logs.Info("req path: %s", string(path))

	req := new(pb.DouyinFeedRequest)
	if err := c.BindAndValidate(&req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	var isLogin bool
	isLogin = false
	token, found := c.Get("token")
	if !found {
		//没登陆
		isLogin = true
	}
	if isLogin == true {
		userInfo, err := authenticate.GetAuthUserInfo(c)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
	}

	resp, err := service.GetFeedInfo(ctx, req, userInfo, isLogin)
	if err != nil {
		logs.Errorf("service err: %v", err)
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
