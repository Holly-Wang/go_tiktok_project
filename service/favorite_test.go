package service

import (
	"fmt"
	"go_tiktok_project/common/dal/mysql"
	pb "go_tiktok_project/idl/pb"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFavOp(t *testing.T) {
	mysql.InitDB()
	var accpet int32 = 1
	//var refuse int32 = 2
	var videoID int64 = 1
	req := new(pb.DouyinFavoriteActionRequest)
	req.ActionType = &accpet
	token, Err := GenerateToken(2, "4")
	assert.NoError(t, Err)
	req.Token = &token
	req.VideoId = &videoID
	res, err := FavoriteAction(req)
	assert.NoError(t, Err)
	fmt.Println(token)
	if err != nil {
		fmt.Println("error:" + err.Error())
		return
	}
	fmt.Println(res)

}

func TestDelOp(t *testing.T) {
	mysql.InitDB()
	var refuse int32 = 2
	var videoID int64 = 1
	req := new(pb.DouyinFavoriteActionRequest)
	req.ActionType = &refuse
	token, Err := GenerateToken(2, "4")
	assert.NoError(t, Err)
	req.Token = &token
	req.VideoId = &videoID
	res, err := FavoriteAction(req)
	assert.NoError(t, Err)
	fmt.Println(res.StatusCode)
	if err != nil {
		fmt.Println("error:" + err.Error())
		return
	}
	fmt.Println(res)

}
