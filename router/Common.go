package router

import (
	jwtAuth "blog-go-api/app/middleware"
	"blog-go-api/app/handle/blog/V1/Common"
	"fmt"
)

/**
 公共不需要鉴权接口
 */
func CommonRouter()  {

	fmt.Println("登陆入口 -- start")
	r := Routers.Group("/api/v1")
	{
		r.POST("/login", jwtAuth.LoginMiddleWare)	// 登陆接口 jwtAuth.LoginMiddleWare
		//r.POST("/logout", jwtAuth.LogoutMiddleWare)	// 登出接口
		r.POST("/getAdminSetting") // 获取配置接口
		r.GET("/getCaptcha", Common.GenerateCaptchaHandler) // 获取验证码
	}
}

