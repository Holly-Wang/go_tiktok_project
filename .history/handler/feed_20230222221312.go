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
	var Token string
	var token string
	isLogin = true
	token = c.Query("token")
	if token == "" {
		//通过token是否为空判断用户是否登录
		isLogin = false
	}
	var userInfo *authenticate.UserInfo
	var info2 
	if isLogin == true {
		Token = token
		info, err := authenticate.CheckToken(Token)
		info2 = info
		//解析token获得用户信息
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			isLogin = false
		}
		else{
			userInfo = info2
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
