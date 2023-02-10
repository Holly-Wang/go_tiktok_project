package service

import (
	"fmt"
	"go_tiktok_project/common"
	"go_tiktok_project/common/dal/mysql"
	"go_tiktok_project/idl/pb"
	"time"

	"gorm.io/gorm"
)

//TODO:雪花主键
func LikeTX(db *gorm.DB, userID int64, videoID int64) error { //点赞视频事务：创造关系并对视频加1
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

func DelTX(db *gorm.DB, userID int64, videoID int64) error { //删除点赞视频事务
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

func FavoriteAction(req *pb.DouyinFavoriteActionRequest) (*pb.DouyinFavoriteActionResponse, error) {
	mysql.InitDB()
	db := mysql.NewDB()
	token := req.Token
	videoID := req.VideoId
	actionType := req.ActionType
	var accpet int32 = 1
	var refuse int32 = 2
	var AC string = "success"
	var WA string = "failes"
	var SuccessCode int32 = 0
	var FailCode int32 = 1
	var (
		rep = new(pb.DouyinFavoriteActionResponse)
	)
	fmt.Println("type:", *actionType)
	if *actionType == accpet {
		userID, err := common.Token2UserID(*token)
		if err != nil {
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: &FailCode,
				StatusMsg:  &WA,
			}, err
		}
		//TODO: 雪花主键
		err = LikeTX(db, int64(userID), *videoID)
		if err != nil {
			fmt.Println("error: " + err.Error())
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: &FailCode,
				StatusMsg:  &WA,
			}, err
		}
		rep.StatusCode = &SuccessCode
		rep.StatusMsg = &AC
	}
	if *actionType == refuse {
		userID, err := common.Token2UserID(*token)
		if err != nil {
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: &FailCode,
				StatusMsg:  &WA,
			}, err
		}
		//先查询视频点赞数，等于0就不用减了
		likeCount, likeErr := mysql.FindLikeOfVideo(*videoID)
		if likeErr != nil {
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: &FailCode,
				StatusMsg:  &WA,
			}, likeErr
		}
		if likeCount == 0 {
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: &SuccessCode,
				StatusMsg:  &AC,
			}, nil
		}
		//点赞数大于0，减一下
		err = DelTX(db, int64(userID), *videoID)
		if err != nil {
			return &pb.DouyinFavoriteActionResponse{
				StatusCode: &FailCode,
				StatusMsg:  &WA,
			}, err
		}
		rep.StatusCode = &SuccessCode
		rep.StatusMsg = &AC
	}
	return rep, nil
}
