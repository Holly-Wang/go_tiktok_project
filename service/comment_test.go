package service

import (
	"fmt"
	model "go_tiktok_project/common/dal/mysql"
	pb_comment "go_tiktok_project/idl/pb_comment"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContCreate(t *testing.T) {
	model.InitDB()
	var create int32 = 1
	var videoID int64 = 1
	var commentid int64 = 5
	var commenttext string = "哈哈哈哈"
	req := new(pb_comment.DouyinCommentActionRequest)
	req.ActionType = &create
	token, Err := GenerateToken(2, "4")
	assert.NoError(t, Err)
	req.Token = &token
	req.VideoId = &videoID
	req.CommentId = &commentid
	req.CommentText = &commenttext
	res, err := CommentActionService(req)
	assert.NoError(t, Err)
	fmt.Println(token)
	if err != nil {
		fmt.Println("error:" + err.Error())
		return
	}
	fmt.Println(res)
}

func TestContDel(t *testing.T) {
	model.InitDB()
	var del int32 = 2
	var videoID int64 = 1
	var commentid int64 = 1
	var commenttext string = "哈哈哈哈"
	req := new(pb_comment.DouyinCommentActionRequest)
	req.ActionType = &del
	token, Err := GenerateToken(2, "4")
	assert.NoError(t, Err)
	req.Token = &token
	req.VideoId = &videoID
	req.CommentId = &commentid
	req.CommentText = &commenttext
	res, err := CommentActionService(req)
	assert.NoError(t, Err)
	fmt.Println(token)
	if err != nil {
		fmt.Println("error:" + err.Error())
		return
	}
	fmt.Println(res)
}
