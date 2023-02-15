package handler

import (
	"context"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"go_tiktok_project/common"
	"go_tiktok_project/common/dal/mysql"
	"go_tiktok_project/common/dal/rediss"
	pb "go_tiktok_project/idl/pb"
	"go_tiktok_project/service"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
)

func UserRegister(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()

	logs.Info("req path: ", path)

	//req := new(pb.DouyinUserRegisterRequest)

	username := c.Query("Username")
	password := c.Query("Password")
	//fmt.Println(username)
	//fmt.Println(password)

	if res := service.IsUsernameLegal(username); !res {
		c.JSON(http.StatusBadRequest, "username is illegal")
		return
	}

	res, err := service.IsUsernameExist(username)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if res {
		c.JSON(http.StatusBadRequest, "username already used")
		return
	}

	userID, err := mysql.CreateUser(username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, "insert to mysql failed")
		return
	}

	token, err := service.GenerateToken(uint64(userID), username)
	if err != nil {
		c.JSON(http.StatusBadRequest, "generate token failed")
		return
	}

	err = rediss.SetToken(ctx, username, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	resp := new(pb.DouyinUserRegisterResponse)

	statuscode := int32(common.RegisterSucces)
	statusmsg := common.RegisterSueecssMsg

	resp.StatusCode = &statuscode
	resp.StatusMsg = &statusmsg
	resp.UserId = &userID
	resp.Token = &token

	c.JSON(http.StatusOK, resp)
}
