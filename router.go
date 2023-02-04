package main

import (
	"go_tiktok_project/common/middlewares"
	handler "go_tiktok_project/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// register registers routers.
func register(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	r.GET("/douyin/user/", middlewares.AuthN(), handler.GetUserInfo)
}
