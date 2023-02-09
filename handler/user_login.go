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
	"reflect"
)

func UserLogin(ctx context.Context, c *app.RequestContext) {
	//defer func() {
	//	recover()
	//}()

	path := c.Request.Path()
	logs.Info("req path: ", path)

	req := new(pb.DouyinUserLoginRequest)
	resp := new(pb.DouyinUserLoginResponse)
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(400, err.Error())
		return
	}

	username := req.GetUsername()
	password := req.GetPassword()

	user, err := mysql.FindUserByNameAndPass(username, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	token, err := rediss.GetTokenByName(ctx, username)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	// return
	*resp.StatusCode = common.LoginSuccess
	*resp.StatusMsg = common.LoginSuccessMsg
	*resp.UserId = cvt2id(user.UserID)
	*resp.Token = token

	c.JSON(http.StatusOK, resp)
}

// cvt2id
func cvt2id(num any) (id int64) {
	return reflect.ValueOf(num).Int()
}
