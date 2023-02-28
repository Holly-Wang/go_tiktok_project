package service

import (
	"go_tiktok_project/common/dal/mysql"
	"testing"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/stretchr/testify/assert"
)

func TestGetUserVideo(t *testing.T) {
	mysql.InitDB()
	list, err := GetUserVideo(123, 111)
	assert.NoError(t, err)
	logs.Info("list: ", list)
}

func TestPostUserVideo(t *testing.T) {
	mysql.InitDB()
	err := PostUserVideo(2, "11", "..//video_data//111//bear//", "bear.mp4")
	assert.NoError(t, err)
}

func TestGetSnapshot(t *testing.T) {
	path, err := GetSnapshot("..//video_data//111//bear//bear.mp4", "..//video_data//111//bear//1", 2)
	assert.NoError(t, err)
	logs.Info("path:", path)

}
