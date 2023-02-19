package handler

import (
	"context"
	pb "go_tiktok_project/idl/biz/model/pb"
	"go_tiktok_project/service"
	"net/http"

	"go_tiktok_project/common/dal/mysql"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/cloudwego/hertz/pkg/app"
)

func Favorite(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()
	logs.Info("req path: %s", string(path))

	req := new(pb.DouyinFavoriteActionRequest)
	if err := c.BindAndValidate(&req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// TODO(liuyiyang): 添加中间件，自动打印接口日志 Req 和 Resp 时延
	resp, err := service.FavoriteAction(req)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func GetFavList(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()
	logs.Info("req path: ", string(path))

	req := new(pb.DouyinFavoriteListRequest)
	if err := c.BindAndValidate(&req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	videoIDs, err := mysql.FindLikeList(int64(req.UserId))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	resp := new(pb.DouyinFavoriteListResponse)
	resp.VideoList = convertFavoriteList(videoIDs)
	c.JSON(http.StatusOK, resp)
}

// TODO(liuyiyang): 完善video信息
func convertFavoriteList(videoIDs []int64) []*pb.Video {
	var ret []*pb.Video
	for _, id := range videoIDs {
		ret = append(ret, &pb.Video{
			Id: id,
		})
	}
	return ret
}
