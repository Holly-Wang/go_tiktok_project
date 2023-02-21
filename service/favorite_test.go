package service

import (
	"fmt"
	"go_tiktok_project/common/authenticate"
	"go_tiktok_project/common/dal/mysql"
	pb "go_tiktok_project/idl/biz/model/pb"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFavOp(t *testing.T) {
	mysql.InitDB()
	req := new(pb.DouyinFavoriteActionRequest)
	req.ActionType = accpet
	token, Err := authenticate.GenToken(2, "4")
	assert.NoError(t, Err)
	req.Token = token
	req.VideoId = 1
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
	req := new(pb.DouyinFavoriteActionRequest)
	req.ActionType = refuse
	token, err := authenticate.GenToken(2, "4")
	assert.NoError(t, err)
	req.Token = token
	req.VideoId = 1
	res, err := FavoriteAction(req)
	assert.NoError(t, err)
	fmt.Println(res.StatusCode)
}
