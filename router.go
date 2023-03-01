package main

import (
	"go_tiktok_project/common/middlewares"
	handler "go_tiktok_project/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// register registers routers.
func register(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	r.GET("/douyin/feed/", handler.GetFeedInfo)

	r.GET("/douyin/favorite/list/", middlewares.AuthN(), handler.GetFavList)
	//r.POST("/douyin/favorite/action/", middlewares.AuthN(), handler.Favorite)
	r.POST("/douyin/favorite/action/", handler.Favorite)

	r.POST("/douyin/user/register/", handler.UserRegister)
	r.POST("/douyin/user/login/", handler.UserLogin)
	r.GET("/douyin/user/", middlewares.AuthN(), handler.GetUserInfo)

	//r.POST("/douyin/comment/action/", middlewares.AuthN(), handler.CommentAction)
	r.POST("/douyin/comment/action/", handler.CommentAction)
	r.GET("/douyin/comment/list/", middlewares.AuthN(), handler.CommentList)

	r.GET("/douyin/publish/list/", middlewares.AuthN(), handler.GetUserVideo)
	//r.POST("/douyin/publish/action/", middlewares.AuthN(), handler.PostUserVideo)
	//r.GET("/douyin/publish/list/", handler.GetUserVideo)
	r.POST("/douyin/publish/action/", handler.PostUserVideo)

}
