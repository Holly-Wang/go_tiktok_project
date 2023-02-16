package service

import (
	"bytes"
	"fmt"
	"go_tiktok_project/common/dal/mysql"
	"os"
	"strings"
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// 生成pb会显示有冲突，就在这里构造了结构体
type User struct {
	Id            int64  `protobuf:"varint,1,req,name=id" json:"id"`                                            // 用户id
	Name          string `protobuf:"bytes,2,req,name=name" json:"name"`                                         // 用户名称
	FollowCount   int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount" json:"follow_count"`       // 关注总数
	FollowerCount int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount" json:"follower_count"` // 粉丝总数
	IsFollow      bool   `protobuf:"varint,5,req,name=is_follow,json=isFollow" json:"is_follow"`                // true-已关注，false-未关注
}

type Video struct {
	Id             int64  `protobuf:"varint,1,req,name=id" json:"id"`                                             // 视频唯一标识
	Author         User   `protobuf:"varint,2,req,name=author" json:"author"`                                     // 视频作者信息
	Play_url       string `protobuf:"bytes,3,req,name=play_url" json:"play_url"`                                  // 视频播放地址
	Cover_url      string `protobuf:"bytes,4,req,name=cover_url_count,json=cover_url" json:"cover_url"`           // 视频封面地址
	Favorite_count int64  `protobuf:"varint,5,req,name=favorite_count,json=favorite_count" json:"favorite_count"` // 视频的点赞总数
	Comment_count  int64  `protobuf:"varint,6,req,name=comment_count,json=comment_count" json:"comment_count"`    // 视频的评论总数
	Is_favorite    bool   `protobuf:"varint,7,req,name=is_favorite ,json=is_favorite" json:"is_favorite "`        // true-已点赞，false-未点赞
	Title          string `protobuf:"bytes,8,req,name=title,json=title" json:"title"`                             // 视频标题
	Abstract       string `protobuf:"bytes,8,req,name=absruct,json=abstract" json:"abstract"`                     // 视频简介
}


// 查询数据库获得用户发布列表
func GetUserVideo(video_user_id, token_user_id int64) ([]Video, error) {

	mysql.InitDB()

	//查询用户信息
	userInfo, err := mysql.FindUserInfoinUser(video_user_id)
	if err != nil {
		logs.Errorf("查询User表出错, error: " + err.Error())
		return nil, err
	}

	//查询token_user是否关注user
	var isFollow bool = false
	FollowCount, err := mysql.FindCountinFollows(token_user_id, video_user_id)
	if err != nil {
		logs.Errorf("查询Follow表出错, error: " + err.Error())
		return nil, err
	}
	if FollowCount > 0 {
		isFollow = true
	}
	user := &User{
		Id:            userInfo.UserID,
		Name:          userInfo.Username,
		FollowCount:   userInfo.Follow_cnt,
		FollowerCount: userInfo.Follower_cnt,
		IsFollow:      isFollow,
	}

	//查询用户发布视频列表
	videos, err := mysql.FindVideoListinVideo(video_user_id)
	if err != nil {
		logs.Errorf("查询Video表出错, error: " + err.Error())
		return nil, err
	}

	var video_list []Video
	for _, v := range videos {
		//查询token_user是否喜欢视频
		var isLike bool = false
		LikesCount, err := mysql.FindCountinLikes(token_user_id, v.VideoID)
		if err != nil {
			logs.Errorf("查询Like表Count出错, error: " + err.Error())
			return nil, err
		}
		if LikesCount > 0 {
			isLike = true
		}

		video := &Video{
			Id:             v.VideoID,
			Author:         *user,
			Play_url:       v.PlayUrl,
			Cover_url:      v.CoverUrl,
			Favorite_count: v.LikeCount,
			Comment_count:  v.CommentCount,
			Is_favorite:    isLike,
			Title:          v.Title,
			Abstract:       v.Abstract,
		}
		video_list = append(video_list, *video)
	}
	
	return video_list, nil
}

// 视频截图，保存视频数据到数据库
func PostUserVideo(user_id int64, title string, filepath string, filedata string) error {

	mysql.InitDB()

	//构造video_id
	max_video_id, err := mysql.FindMaxIdinVideos()
	if err != nil {
		logs.Errorf("查询Video表主键出错, error: " + err.Error())
		return err
	}
	var video_id int64 = max_video_id + 1

	//将视频第1帧截图保存为封面
	strArray := strings.Split(filepath, ".")
	ImageName := strArray[0]
	logs.Info("imageName: %s", ImageName)
	logs.Info("filepath: %s", filepath)
	imagePath, err := GetSnapshot(filepath, ImageName, 1)
	if err != nil {
		logs.Errorf("截取视频第一帧错误, error: " + err.Error())
		return err
	}
	logs.Info("imagePath: ", imagePath)

	//将video数据写入数据库
	err_createvideo := mysql.CreateVideo(video_id, user_id, filepath, imagePath, 0, 0, title, filedata)
	if err_createvideo != nil {
		logs.Errorf("创建Video数据出错, error: " + err.Error())
		return err
	}
	return nil
}

// 截取视频为封面
func GetSnapshot(videoPath, imageName string, frameNum int) (ImagePath string, err error) {

	//截取视频
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		logs.Errorf("截图失败, error: ", err)
		return "", err
	}

	//保存为图片
	img, err := imaging.Decode(buf)
	if err != nil {
		logs.Info("生成缩略图失败：", err)
		return "", err
	}
	err = imaging.Save(img, imageName+".png")
	if err != nil {
		logs.Info("生成缩略图失败：", err)
		return "", err
	}

	imgPath := imageName + ".png"

	return imgPath, nil
}
