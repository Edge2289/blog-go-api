package router

import (
	"blog-go-api/app/handle/blog/V1/Common"
	"blog-go-api/app/middleware"
	jwtAuth "blog-go-api/app/middleware"
	"fmt"
)

/**
 公共不需要鉴权接口
 */
func CommonRouter()  {

	fmt.Print("登陆入口加载")
	// 这里都是一些不需要登录的接口
	r := Routers.Group(apiRouter + "v1" )
	{
		// 登陆 登出
		r.POST("/login", jwtAuth.LoginMiddleWare)	// 登陆接口 jwtAuth.LoginMiddleWare
		//r.POST("/logout", jwtAuth.LogoutMiddleWare)	// 登出接口

		r.POST("/getAdminSetting") // 获取配置接口
		r.GET("/getCaptcha", Common.GenerateCaptchaHandler) // 获取验证码

		// 添加访问日志
		r.GET("/addAccess", Common.AddAccess)
	}

	// 公共接口加载
	r.Use(middleware.CheckJwt())
	{
		// 菜单管理  Menu
		//r.POST()

		// 日志管理  Log

		// 系统配置  Sys

		// 访问管理	Access
	}
}

