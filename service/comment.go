package service

import (
	"github.com/cloudwego/hertz/cmd/hz/util/logs"
	"go_tiktok_project/common/authenticate"
	model "go_tiktok_project/common/dal/mysql"
	"go_tiktok_project/common/middlewares"
	"go_tiktok_project/idl/biz/model/pb"
)

const (
	cre int32 = 1
	del int32 = 2

	AC string = "success"
	WA string = "failes"

	SuccessCode int32 = 0
	FailCode    int32 = 1
)

func CommentActionService(req *pb.DouyinCommentActionRequest) (*pb.DouyinCommentActionResponse, error) {
	var (
		resp         = new(pb.DouyinCommentActionResponse)
		comment      pb.Comment
		comment_user pb.User
	)
	userInfo, err := authenticate.CheckToken(req.Token)
	if err != nil {
		// 没有调用过auth
		middlewares.AuthN()
	}
	userID := userInfo.UserID
	switch req.ActionType {
	case cre:
		logs.Info("%d", userID)
		if err != nil {
			return &pb.DouyinCommentActionResponse{
				StatusCode: FailCode,
				StatusMsg:  WA,
			}, err
		}
		user, err := model.FindUserById(uint64(userID))
		if err != nil {
			return &pb.DouyinCommentActionResponse{
				StatusCode: FailCode,
				StatusMsg:  WA,
			}, err
		}
		userrID := model.FindVidByUid(req.VideoId)
		is_follow, err := model.CheckFollow(userID, userrID)
		if err != nil {
			logs.Errorf("[SQL Error] check follow err: %v", err)
			return nil, err
		}

		userName := user.Username
		follow_count := user.Follow_cnt
		follower_count := user.Follower_cnt
		userID_int64 := int64(userID)
		comment_user.Id = userID_int64
		comment_user.Name = userName
		comment_user.FollowCount = follow_count
		comment_user.FollowerCount = follower_count
		comment_user.IsFollow = is_follow

		comment.Id = req.CommentId
		comment.Content = req.CommentText
		comment.User = &comment_user

		// 创建时默认自己没有点赞且点赞数为0
		model.CreateComment(req.CommentId, req.VideoId, userID_int64, req.CommentText, 0, false)

		return &pb.DouyinCommentActionResponse{
			StatusCode: SuccessCode,
			StatusMsg:  AC,
			Comment:    &comment,
		}, err
	}
	if req.ActionType == del {
		user, err := model.FindUserById(uint64(userID))
		if err != nil {
			return &pb.DouyinCommentActionResponse{
				StatusCode: FailCode,
				StatusMsg:  WA,
			}, err
		}
		userrID := model.FindVidByUid(req.VideoId)
		is_follow, err := model.CheckFollow(int64(userID), userrID)
		if err != nil {
			logs.Errorf("[SQL Error] check follow err: %v", err)
			return nil, err
		}
		userName := user.Username
		follow_count := user.Follow_cnt
		follower_count := user.Follower_cnt
		userID_int64 := int64(userID)

		comment_user.Id = userID_int64
		comment_user.Name = userName
		comment_user.FollowCount = follow_count
		comment_user.FollowerCount = follower_count
		comment_user.IsFollow = is_follow

		comment.Id = req.CommentId
		comment.User = &comment_user

		model.DelComment(req.CommentId)
		return &pb.DouyinCommentActionResponse{
			StatusCode: SuccessCode,
			StatusMsg:  AC,
			Comment:    &comment,
		}, err
	}
	return resp, nil
}
