package handler

import (
	"context"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/cloudwego/hertz/pkg/app"
	"go_tiktok_project/common"
	"go_tiktok_project/common/dal/mysql"
	"go_tiktok_project/common/dal/rediss"
	pb "go_tiktok_project/idl/pb"
	"net/http"
)

func UserLogin(ctx context.Context, c *app.RequestContext) {
	//defer func() {
	//	recover()
	//}()

	path := c.Request.Path()
	logs.Info("req path: ", path)

	resp := new(pb.DouyinUserLoginResponse)

	username := c.Query("Username")
	password := c.Query("Password")

	user, err := mysql.FindUserByNameAndPass(username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := rediss.GetTokenByName(ctx, username)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// return

	loginsuccode := int32(common.LoginSuccess)
	loginsucmsg := common.LoginSuccessMsg

	resp.StatusCode = &loginsuccode
	resp.StatusMsg = &loginsucmsg
	resp.UserId = &user.UserID
	resp.Token = &token

	c.JSON(http.StatusOK, resp)
}
