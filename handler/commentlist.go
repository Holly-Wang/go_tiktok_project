package handler

import (
	"context"
	"net/http"
	"strconv"

	model "go_tiktok_project/common/dal/mysql"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/cloudwego/hertz/pkg/app"
)

type CommentListResponse struct {
	StatusCode  int32           `json:"status_code"`
	StatusMsg   string          `json:"status_msg,omitempty"`
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

// 查看视频的所有评论，按发布时间倒序
func CommentList(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()
	logs.Info("req path: %s", path)

	model.InitDB()

	// token := c.Query("token")
	Video_id := c.Query("video_id")
	video_id_n, _ := strconv.Atoi(Video_id)
	video_id := int64(video_id_n)
	list, err := model.FindComment(video_id)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	resp := new(CommentListResponse)
	resp.StatusCode = 0
	resp.StatusMsg = "Success"
	resp.CommentList = list
	c.JSON(http.StatusOK, resp)
}
