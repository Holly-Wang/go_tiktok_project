package service

import (
	"bytes"
	"fmt"
	"go_tiktok_project/common/dal/mysql"
	"go_tiktok_project/idl/biz/model/pb"
	"os"
	"strings"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

// 查询数据库获得用户发布列表
func GetUserVideo(videoUserID, tokenUserID int64) ([]*pb.Video, error) {
	//查询用户信息
	userInfo, err := mysql.FindUserInfoinUser(videoUserID)
	if err != nil {
		logs.Errorf("查询User表出错, error: " + err.Error())
		return nil, err
	}

	//查询token_user是否关注user
	var isFollow bool = false
	FollowCount, err := mysql.FindCountinFollows(tokenUserID, videoUserID)
	if err != nil {
		logs.Errorf("查询Follow表出错, error: " + err.Error())
		return nil, err
	}
	if FollowCount > 0 {
		isFollow = true
	}
	user := &pb.User{
		Id:            userInfo.UserID,
		Name:          userInfo.Username,
		FollowCount:   userInfo.Follow_cnt,
		FollowerCount: userInfo.Follower_cnt,
		IsFollow:      isFollow,
	}

	//查询用户发布视频列表
	videos, err := mysql.FindVideoListinVideo(videoUserID)
	if err != nil {
		logs.Errorf("查询Video表出错, error: " + err.Error())
		return nil, err
	}

	var videoList []*pb.Video
	for _, v := range videos {
		//查询token_user是否喜欢视频
		var isLike bool = false
		LikesCount, err := mysql.FindCountinLikes(tokenUserID, v.VideoID)
		if err != nil {
			logs.Errorf("查询Like表Count出错, error: " + err.Error())
			return nil, err
		}
		if LikesCount > 0 {
			isLike = true
		}

		video := &pb.Video{
			Id:            v.VideoID,
			Author:        user,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.LikeCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    isLike,
			Title:         v.Title,
		}
		videoList = append(videoList, video)
	}

	return videoList, nil
}

// 视频截图，保存视频数据到数据库
func PostUserVideo(user_id int64, title string, filedir, filename string) error {

	mysql.InitDB()

	//构造video_id
	max_video_id, err := mysql.FindMaxIdinVideos()
	if err != nil {
		logs.Errorf("查询Video表主键出错, error: " + err.Error())
		return err
	}
	var video_id int64 = max_video_id + 1

	//将视频第1帧截图保存为封面
	strArray := strings.Split(filename, ".")
	filedata := strArray[0]
	ImageName := filedir + filedata
	filepath := filedir + filename
	logs.Info("imageName: %s", ImageName)
	logs.Info("filepath: %s", filepath)
	imagePath, err := GetSnapshot(filepath, ImageName, 1)
	if err != nil {
		logs.Errorf("截取视频第一帧错误, error: " + err.Error())
		return err
	}
	logs.Info("imagePath: ", imagePath)

	save_filename := fmt.Sprintf("%d_%d_%s", user_id, video_id, filedata)
	video_path := "/root/plalyy/go/src/go_tiktok_project/video_data/" + save_filename + ".mp4"
	image_path := "/root/plalyy/go/src/go_tiktok_project/video_data/" + save_filename + ".jpg"

	err = SaveVideoToService(filepath, video_path)
	if err != nil {
		logs.Errorf("上传视频出错, error: " + err.Error())
		return err
	}
	err = SaveVideoToService(imagePath, image_path)
	if err != nil {
		logs.Errorf("上传图片出错, error: " + err.Error())
		return err
	}

	video_database_path := "http://49.232.155.203:8807/data/" + save_filename + ".mp4"
	image_database_path := "http://49.232.155.203:8807/data/" + save_filename + ".jpg"

	logs.Info("video_database_path:", video_database_path)
	logs.Info("image_database_path", image_database_path)
	//将video数据写入数据库
	err_createvideo := mysql.CreateVideo(video_id, user_id, video_database_path, image_database_path, 0, 0, title, filedata)
	if err_createvideo != nil {
		logs.Errorf("创建Video数据出错, error: " + err.Error())
		return err
	}
	return nil
}

// 保存视频到服务器
func SaveVideoToService(localFile, RemoteFileName string) (err error) {
	username := "root"
	password := "Freedom9"
	ip := "49.232.155.203"
	port := "22"
	client := NewSSHClient(username, password, ip, port)

	n, err := client.UploadFile(localFile, RemoteFileName)
	if err != nil {
		logs.Errorf("上传到服务器失败, error: ", err)
		return err
	}
	logs.Info("upload file[%v] ok, size=[%d]\n", RemoteFileName, n)
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
