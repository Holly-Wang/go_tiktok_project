package handler

import (
	"context"
	pb "go_tiktok_project/idl/biz/model/pb"
	"go_tiktok_project/service"
	"net/http"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"

	"github.com/cloudwego/hertz/pkg/app"
)

func UserRegister(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()
	logs.Info("req path: %s", string(path))

	req := new(pb.DouyinUserRegisterRequest)
	if err := c.BindAndValidate(&req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	logs.Infof("req: %s, %s", req.Username, req.Password)

	resp, err := service.UserRegister(ctx, req)
	if err != nil {
		// TODO(liuyiyang): 封装 err，在service中体现出不同的http返回值, 并且修改 resp 的 code 和 msg
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
