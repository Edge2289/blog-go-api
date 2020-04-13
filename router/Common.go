package router

import (
	"blog-go-api/app/handle/blog/V1/Admin"
	//jwtAuth "blog-go-api/app/middleware"
	"fmt"
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
	fmt.Println("登陆入口")
	r := Routers.Group("/api/v1")
	{
		r.POST("/login", Admin.Login)	// 登陆接口
		r.POST("/logout", )	// 登出接口
		r.POST("/getAdminSetting") // 获取配置接口
	}
}

