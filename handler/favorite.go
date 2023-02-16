package handler

import (
	"context"
	pb "go_tiktok_project/idl/pb"
	"go_tiktok_project/service"
	"strconv"

	"go_tiktok_project/common/dal/mysql"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/cloudwego/hertz/pkg/app"
)

//点赞操作的response在favorite.pb.go里生成了，DouyinFavoriteActionResponse
//点赞列表
type FavListOpResp struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	VideoList  interface{} `json:"video_list"`
}

type FavListReq struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func Favorite(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()
	logs.Info("req path: %s", path)

	req := new(pb.DouyinFavoriteActionRequest)
	ActionType := c.Query("action_type")
	actionType_n, _ := strconv.Atoi(ActionType)
	actionType := int32(actionType_n)
	req.ActionType = &actionType
	vidio_id := c.Query("video_id")
	vidioid, _ := strconv.Atoi(vidio_id)
	video_id := int64(vidioid)
	req.VideoId = &video_id
	token := c.Query("token")
	req.Token = &token

	//校验并解析请求
	// if err := c.BindAndValidate(&req); err != nil {
	// 	c.String(400, err.Error())
	// 	return
	// }
	logs.Info("req actiontype:%v", req.ActionType)
	resp, err := service.FavoriteAction(req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(200, resp)
}

func GetFavList(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()
	logs.Info("req path: %s", path)
	// if err := c.BindAndValidate(&req); err != nil {
	// 	c.String(400, err.Error())
	// 	return
	// }
	UserID := c.Query("user_id")
	userID, _ := strconv.Atoi(UserID)
	lists, err := mysql.FindLikeList(int64(userID))
	if err != nil {
		c.String(400, err.Error())
		return
	}
	resp := new(FavListOpResp)
	resp.VideoList = lists
	resp.StatusCode = 0
	resp.StatusMsg = "success"
	c.JSON(200, resp)
}
