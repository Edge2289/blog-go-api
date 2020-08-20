package router

import (
	"blog-go-api/app/middleware"
)

func SystemRouter(base string) {
	r := Routers.Group("api/v1/" + base)
	r.Use(middleware.CheckJwt())
	{
		// 菜单管理
		//r.POST()

		// 日志管理

		// 系统配置

		// 访问管理
	}
}
