package main

import (
	"go_tiktok_project/common/middlewares"
	"go_tiktok_project/common/middlewares/TokenCheck"
	handler "go_tiktok_project/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// register registers routers.
func register(r *server.Hertz) {
	r.GET("/ping", handler.Ping)
	r.GET("/douyin/feed/", TokenCheck.TokenCheck(), handler.GetFeedInfo)
	r.GET("/douyin/user/", middlewares.AuthN(), handler.GetUserInfo)
	r.GET("/douyin/favorite/list/", handler.GetFavList)

	r.POST("/douyin/user/register", handler.UserRegister)
	r.POST("/douyin/user/login", handler.UserLogin)

	r.POST("/douyin/comment/action", handler.CommentAction)
	r.GET("/douyin/comment/list", handler.CommentList)
	r.POST("/douyin/favorite/action/", handler.Favorite)

	r.GET("/douyin/publish/list/", handler.GetUserVideo)
	r.POST("/douyin/publish/action/", handler.PostUserVideo)
}
