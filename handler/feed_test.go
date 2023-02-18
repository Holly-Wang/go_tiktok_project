package handler

import (
	"context"
	"fmt"
	"go_tiktok_project/common/authenticate"
	"go_tiktok_project/common/middlewares"
	pb "go_tiktok_project/idl/pb"
	"go_tiktok_project/service"
	"testing"
)

func TestGetFeedInfo(t *testing.T) {
	var DF = new(pb.DouyinFeedRequest)
	var User = new(authenticate.UserInfo)
	var name string
	var ps string
	var tokenString string
	name = "4"
	ps = "5"
	tokenString, _ = middlewares.GenToken(name, ps)
	fmt.Println(tokenString)
	DF.Token = &tokenString
	//fmt.Println(*DF.Token)
	User.Username = "2"
	//var rep *pb.DouyinFeedResponse
	fmt.Println(service.GetFeedInfo(context.Background(), DF, User))
}
