package handler

import (
	"context"
	"fmt"
	"go_tiktok_project/common/authenticate"
	pb "go_tiktok_project/idl/biz/model/pb"
	"go_tiktok_project/service"
	"testing"
)

func TestGetFeedInfo(t *testing.T) {
	var req = new(pb.DouyinFeedRequest)
	var User = new(authenticate.UserInfo)
	var tokenString string
	name := "4"
	tokenString, _ = authenticate.GenToken(1, name)
	req.Token = tokenString
	User.Username = "2"
	fmt.Println(service.GetFeedInfo(context.Background(), req, User))
}
