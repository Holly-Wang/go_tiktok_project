// Code generated by hertz generator.

package handler

import (
	"context"
	"fmt"
	"go_tiktok_project/common"
	"go_tiktok_project/common/authenticate"
	"go_tiktok_project/common/middlewares"
	"go_tiktok_project/idl/biz/model/pb"
	"go_tiktok_project/service"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// 获得发布列表
func GetUserVideo(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()
	logs.Info("req path: %s", string(path))

	req := new(pb.DouyinPublishListRequest)
	if err := c.BindAndValidate(&req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	userInfo, err := authenticate.GetAuthUserInfo(c)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	//获取用户发布列表
	video_list, err := service.GetUserVideo(req.UserId, userInfo.UserID)
	if err != nil {
		logs.Errorf("service err, error: " + err.Error())
		c.JSON(http.StatusBadRequest, pb.DouyinPublishListResponse{
			StatusCode: common.GetUserVideoFailed,
			StatusMsg:  common.GetUserVideoFailedMsg,
		})
		return
	}

	resp := &pb.DouyinPublishListResponse{
		StatusCode: common.GetUserVideoSuccess,
		StatusMsg:  common.GetUserVideoSuccessMsg,
		VideoList:  video_list,
	}
	c.JSON(consts.StatusOK, resp)
}

// 用户视频投稿
func PostUserVideo(ctx context.Context, c *app.RequestContext) {
	path := c.Request.Path()
	logs.Info("req path: %s", string(path))

	//userInfo, err := authenticate.GetAuthUserInfo(c)
	//if err != nil {
	//	c.String(http.StatusBadRequest, err.Error())
	//	return
	//}

	userInfo, err := authenticate.CheckToken(c.PostForm("token"))
	if err != nil {
		// 没有调用过auth
		middlewares.AuthN()
	}

	title, _ := c.GetPostForm("title")
	data, err := c.FormFile("data")
	if err != nil {
		panic(err)
	}

	//利用文件名获得文件信息和文件路径
	filename := filepath.Base(data.Filename)
	fileInfo := strings.Split(filename, ".")
	filedir := fileInfo[0]
	filedir = fmt.Sprintf("./video_data/%d/%s/", userInfo.UserID, filedir) //token是没解析的

	//创建存储文件夹
	_, erByStat := os.Stat(filedir)
	if erByStat != nil {
		logs.Errorf("os stat %s error......%v", filedir, erByStat)
	}
	//该判断主要是部分文件权限问题导致os.Stat()出错,具体看业务启用
	//使用os.IsNotExist()判断为true,说明文件夹不存在
	if os.IsNotExist(erByStat) {
		logs.Info("%s is not exist", erByStat.Error())
		err := os.MkdirAll(filedir, 0777)
		if err != nil {
			logs.Error("创建文件夹错误 , err: %v", err)
			c.JSON(http.StatusBadRequest, utils.H{
				"status_code": common.PublishFailed,
				"status_msg":  common.PublishFailedMsg,
			})
			return
		} else {
			logs.Info("Create dir %s success!", filedir)
		}
	}

	//保存视频文件

	saveFile := filedir + filename
	logs.Info("filepath: %s", saveFile)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		logs.Error("保存视频失败, err: %v", err)
		c.JSON(http.StatusBadRequest, utils.H{
			"status_code": common.PublishFailed,
			"status_msg":  common.PublishFailedMsg,
		})
		return
	}

	//service 保存视频数据到数据库
	err = service.PostUserVideo(userInfo.UserID, title, filedir, filename)
	if err != nil {
		logs.Error("server error, err: %v", err)
		c.JSON(http.StatusBadRequest, utils.H{
			"status_code": common.PublishFailed,
			"status_msg":  common.PublishFailedMsg,
		})
		return
	}

	//返回参数
	c.JSON(consts.StatusOK, utils.H{
		"status_code": common.PublishSuccess,
		"status_msg":  common.PublishSuccessMsg,
	})
}
