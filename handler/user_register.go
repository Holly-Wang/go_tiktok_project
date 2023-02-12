package handler

import (
	"context"
	"go_tiktok_project/common"
	"go_tiktok_project/common/dal/mysql"
	"go_tiktok_project/common/dal/rediss"
	pb "go_tiktok_project/idl/pb"
	"go_tiktok_project/service"
	"net/http"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/cloudwego/hertz/pkg/app"
)

func UserRegister(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()

	logs.Info("req path: ", path)

	req := new(pb.DouyinUserRegisterRequest)

	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if res := service.IsUsernameLegal(req.GetUsername()); !res {
		c.JSON(http.StatusBadRequest, "username is illegal")
		return
	}

	res, err := service.IsUsernameExist(req.GetUsername())
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if res {
		c.JSON(http.StatusBadRequest, "username already used")
		return
	}

	userID, err := mysql.CreateUser(req.GetUsername(), req.GetPassword())
	if err != nil {
		c.JSON(http.StatusBadRequest, "insert to mysql failed")
		return
	}

	token, err := service.GenerateToken(uint64(userID), req.GetUsername())
	if err != nil {
		c.JSON(http.StatusBadRequest, "generate token failed")
		return
	}

	err = rediss.SetToken(ctx, req.GetUsername(), token)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resp := new(pb.DouyinUserRegisterResponse)

	*resp.StatusCode = common.LoginSuccess
	*resp.StatusMsg = "register success"
	// userid uint64 or int64?
	*resp.UserId = cvt2id(userID)
	*resp.Token = token

	c.JSON(http.StatusOK, resp)
}
