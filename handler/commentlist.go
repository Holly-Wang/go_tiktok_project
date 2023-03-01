package handler

import (
	"context"
	"net/http"

	model "go_tiktok_project/common/dal/mysql"
	"go_tiktok_project/idl/biz/model/pb"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/cloudwego/hertz/pkg/app"
)

// CommentList 查看视频的所有评论，按发布时间倒序
func CommentList(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()
	logs.Info("req path: %s", string(path))

	req := new(pb.DouyinCommentListRequest)
	if err := c.BindAndValidate(&req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	list, err := model.FindComment(req.VideoId)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp := new(pb.DouyinCommentListResponse)
	resp.CommentList = convertComment(list)
	c.JSON(http.StatusOK, resp)
}

func convertComment(comments []model.Comment) []*pb.Comment {
	var ret []*pb.Comment
	for _, v := range comments {
		ret = append(ret, &pb.Comment{
			Id: v.CommentID,
			User: &pb.User{
				Id: v.UserID,
			},
			Content:    v.Context,
			CreateDate: v.CommentTime.Format("2006-01-02 15:04:05"),
		})
	}
	return ret
}
