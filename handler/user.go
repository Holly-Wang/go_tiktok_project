package handler

import (
	"context"
	pb "go_tiktok_project/idl/pb"

	"github.com/cloudwego/hertz/pkg/app"
)

func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req pb.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := new(pb.DouyinUserResponse)

	resp.StatusCode = new(int32)

	c.JSON(200, resp)
}
