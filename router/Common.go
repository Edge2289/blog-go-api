package router

import (
	//jwtAuth "blog-go-api/app/middleware"
	//"fmt"
)

/**
 公共不需要鉴权接口
 */
func CommonRouter()  {

	// the jwt middleware
	//authMiddleware, err := jwtAuth.AuthInit()
	//
	//if err != nil {
	//	_ = fmt.Errorf("JWT Error", err.Error())
	//}
	//r := Routers.Group("/api/v1")
	//{
	//	r.POST("/login", authMiddleware.LoginHandler)	// 登陆接口
	//	r.POST("/logout", )	// 登出接口
	//	r.POST("/getAdminSetting") // 获取配置接口
	//}
}

