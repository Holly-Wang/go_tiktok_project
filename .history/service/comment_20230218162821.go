package service

import (
	common "go_tiktok_project/common"
	model "go_tiktok_project/common/dal/mysql"
	pb_comment "go_tiktok_project/idl/pb_comment"
	"strconv"

	"github.com/cloudwego/hertz/cmd/hz/util/logs"
)

func CommentActionService(req *pb_comment.DouyinCommentActionRequest) (*pb_comment.DouyinCommentActionResponse, error) {
	model.InitDB()
	var cre int32 = 1
	var del int32 = 2
	var AC string = "success"
	var WA string = "failes"
	var SuccessCode int32 = 0
	var FailCode int32 = 1
	var ZERO int64 = 0
	token := req.Token
	videoid := req.VideoId
	actiontype := req.ActionType
	commentid := req.CommentId
	commenttext := req.CommentText
	var (
		rsp = new(pb_comment.DouyinCommentActionResponse)
	)
	var comment pb_comment.Comment
	var comment_user pb_comment.User

	if *actiontype == cre {

		userID, err := common.Token2UserID(*token)
		logs.Info(strconv.FormatUint(userID, 10))
		if err != nil {
			return &pb_comment.DouyinCommentActionResponse{
				StatusCode: &FailCode,
				StatusMsg:  &WA,
			}, err
		}
		user, err := model.FindUserById(userID)
		if err != nil {
			return &pb_comment.DouyinCommentActionResponse{
				StatusCode: &FailCode,
				StatusMsg:  &WA,
			}, err
		}
		userName := user.Username
		follow_count := user.Follow_cnt
		follower_count := user.Follower_cnt
		userID_int64 := int64(userID)
		// todo 加一个 is_follow【合并后进行】

		comment_user.Id = &userID_int64
		comment_user.Name = &userName
		comment_user.FollowCount = &follow_count
		comment_user.FollowerCount = &follower_count

		comment.Id = commentid
		comment.Content = commenttext

		comment.User = &comment_user

		// 创建时默认自己没有点赞且点赞数为0
		model.CreateComment(*commentid, *videoid, userID_int64, *commenttext, ZERO, false)

		return &pb_comment.DouyinCommentActionResponse{
			StatusCode: &SuccessCode,
			StatusMsg:  &AC,
			Comment:    &comment,
		}, err
	}
	if *actiontype == del {

		userID, err := common.Token2UserID(*token)
		if err != nil {
			return &pb_comment.DouyinCommentActionResponse{
				StatusCode: &FailCode,
				StatusMsg:  &WA,
			}, err
		}
		user, err := model.FindUserById(userID)
		if err != nil {
			return &pb_comment.DouyinCommentActionResponse{
				StatusCode: &FailCode,
				StatusMsg:  &WA,
			}, err
		}
		userName := user.Username
		follow_count := user.Follow_cnt
		follower_count := user.Follower_cnt
		userID_int64 := int64(userID)
		// todo 加一个 is_follow【合并后进行】

		comment_user.Id = &userID_int64
		comment_user.Name = &userName
		comment_user.FollowCount = &follow_count
		comment_user.FollowerCount = &follower_count

		comment.Id = commentid

		comment.User = &comment_user

		model.DelComment(*commentid)
		return &pb_comment.DouyinCommentActionResponse{
			StatusCode: &SuccessCode,
			StatusMsg:  &AC,
			Comment:    &comment,
		}, err
	}

	return rsp, nil

}
