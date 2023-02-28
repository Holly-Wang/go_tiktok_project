package service

import (
	"go_tiktok_project/common/authenticate"
	"go_tiktok_project/common/dal/mysql"
	"go_tiktok_project/common/middlewares"
	"go_tiktok_project/idl/biz/model/pb"
	"time"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"gorm.io/gorm"
)

// TODO:雪花主键
// 点赞视频事务：创造关系并对视频加1
func LikeTX(db *gorm.DB, userID int64, videoID int64) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	like := mysql.Like{OwnerID: userID, VideoID: videoID, LikeTime: time.Now()}
	err := tx.Create(&like).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Model(&mysql.Video{}).Where("video_id=?", videoID).Update("like_count", gorm.Expr("like_count+?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

// 删除点赞视频事务
func DelTX(db *gorm.DB, userID int64, videoID int64) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	var like mysql.Like
	err := tx.Where("owner_id = ? AND video_id = ?", userID, videoID).First(&like).Error
	if err != nil {
		return err
	}
	err = tx.Model(&mysql.Video{}).Where("video_id=?", videoID).Update("like_count", gorm.Expr("like_count+?", -1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	keyID := like.KeyID
	like = mysql.Like{KeyID: keyID}
	err = tx.Delete(&like).Error
	if err != nil {
		return err
	}
	return tx.Commit().Error
}

const (
	accpet int32 = 1
	refuse int32 = 2
)

func FavoriteAction(req *pb.DouyinFavoriteActionRequest) (*pb.DouyinFavoriteActionResponse, error) {
	db := mysql.NewDB()
	var WA string = "failes"
	var FailCode int32 = 1
	resp := new(pb.DouyinFavoriteActionResponse)

	userInfo, err := authenticate.CheckToken(req.Token)
	if err != nil {
		// 没有调用过auth
		middlewares.AuthN()
	}

	userID := userInfo.UserID
	switch req.ActionType {
	case accpet:
		err := LikeTX(db, int64(userID), req.VideoId)
		if err != nil {
			logs.Error("err: %v", err)
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: FailCode,
				StatusMsg:  WA,
			}, err
		}
	case refuse:
		// 先查询视频点赞数，等于0就不用减了
		likeCount, likeErr := mysql.FindLikeOfVideo(req.VideoId)
		if likeErr != nil {
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: FailCode,
				StatusMsg:  WA,
			}, likeErr
		}
		if likeCount == 0 {
			// TODO(liuyiyang): 不赋值是不是就是空的
			return &pb.DouyinFavoriteActionResponse{}, nil
		}
		// 点赞数大于0，减一下
		err := DelTX(db, int64(userID), req.VideoId)
		if err != nil {
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: FailCode,
				StatusMsg:  WA,
			}, err
		}
	}
	return resp, nil
}
