package handler

import (
	"context"
	pb "go_tiktok_project/idl/pb"
	service "go_tiktok_project/service"
	"net/http"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/cloudwego/hertz/pkg/app"
)

func UserLogin(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()
	logs.Info("req path: ", string(path))

	req := new(pb.DouyinUserLoginRequest)
	if err := c.BindAndValidate(&req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := service.UserLogin(ctx, req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
