package router

import (
	"blog-go-api/app/handle/blog/V1/Common"
	jwtAuth "blog-go-api/app/middleware"
)

/**
公共不需要鉴权接口
*/
func CommonRouter() {
	/**
		需要鉴权的
	 */
	RoutersAuth := Routers.Group(apiRouter + "v1")
	RoutersAuth.Use(jwtAuth.CheckJwt())

	// 这里都是一些不需要登录的接口
	r := Routers.Group(apiRouter + "v1")
	{
		// 登陆 登出
		r.POST("/login", jwtAuth.LoginMiddleWare) // 登陆接口 jwtAuth.LoginMiddleWare
		//r.POST("/logout", jwtAuth.LogoutMiddleWare)	// 登出接口

		r.POST("/getAdminSetting")                          // 获取配置接口
		r.GET("/getCaptcha", Common.GenerateCaptchaHandler) // 获取验证码

		// 添加访问日志
		r.GET("/addAccess", Common.AddAccess)
	}
	r.Group(apiRouter + "v1")
	{
		// 上传图片
		RoutersAuth.GET("/uploadImg", Common.GetImg)
		RoutersAuth.POST("/uploadImg", Common.UploadImg)
		RoutersAuth.DELETE("/uploadImg", Common.DelImg)
		RoutersAuth.PUT("/uploadImg", Common.MvImgCategory)

		// 图片分组管理
		RoutersAuth.GET("/imgGroup", Common.GetImgGroup)
		RoutersAuth.DELETE("/imgGroup", Common.DelImgGroup)
		RoutersAuth.PUT("/imgGroup", Common.UpdateImgGroup)
		RoutersAuth.POST("/imgGroup", Common.AddImgGroup)
	}

	// 公共接口加载
	r.Group(apiRouter + "v1")
	{
		// 上传图片
		//r.POST("/uploadImgs", Common.Upload)
		//r.DELETE("/uploadImgs", Common.Delete)

		// 菜单管理  Menu
		//r.POST()

		// 日志管理  Log

		// 系统配置  Sys

		// 访问管理	Access
	}
}
