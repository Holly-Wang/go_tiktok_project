package handler

import (
	"context"
	pb_comment "go_tiktok_project/idl/pb_comment"
	service "go_tiktok_project/service"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/cloudwego/hertz/pkg/app"
)

// 登录用户对视频进行评论
func CommentAction(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()
	logs.Info("req path: %s", path)

	req := new(pb_comment.DouyinCommentActionRequest)

	token := c.Query("token")
	req.Token = &token
	Video_id := c.Query("video_id")
	video_id_n, _ := strconv.Atoi(Video_id)
	video_id := int64(video_id_n)
	req.VideoId = &video_id
	ActionType := c.Query("action_type")
	actionType_n, _ := strconv.Atoi(ActionType)
	actionType := int32(actionType_n)
	req.ActionType = &actionType
	comment_text := c.Query("comment_text")
	req.CommentText = &comment_text
	Comment_id := c.Query("comment_id")
	comment_id_n, _ := strconv.Atoi(Comment_id)
	comment_id := int64(comment_id_n)
	req.VideoId = &comment_id

	resp, err := service.CommentActionService(req)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)

}
