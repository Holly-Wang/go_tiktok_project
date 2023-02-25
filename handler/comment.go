package handler

import (
	"context"
	"go_tiktok_project/idl/biz/model/pb"
	service "go_tiktok_project/service"
	"net/http"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/cloudwego/hertz/pkg/app"
)

// CommentAction 登录用户对视频进行评论
func CommentAction(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()
	logs.Info("req path: %s", string(path))

	req := new(pb.DouyinCommentActionRequest)
	if err := c.BindAndValidate(&req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp, err := service.CommentActionService(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}
