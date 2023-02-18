package handler

import (
	"context"
	"fmt"
	"go_tiktok_project/common/authenticate"
	"go_tiktok_project/common/middlewares/TokenCheck"
	"go_tiktok_project/idl/pb_feed"
	"go_tiktok_project/service"
	"testing"
)

func TestGetFeedInfo(t *testing.T) {
	var req = new(pb_feed.DouyinFeedRequest)
	var User = new(authenticate.UserInfo)
	var tokenString string
	name := "4"
	ps := "5"
	tokenString, _ = TokenCheck.GenToken(name, ps)
	fmt.Println(tokenString)
	req.Token = &tokenString
	User.Username = "2"
	fmt.Println(service.GetFeedInfo(context.Background(), req, User))
}
