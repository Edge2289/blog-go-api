package router

import (
	"blog-go-api/app/middleware"
)

func ArticleRouter(base string)  {

	r := Routers.Group("/api/v1/" + base)

	//r.Use(middleware.CheckJwt())
	//{
		r.GET("/deptList", middleware.CheckJwt)
	//}
}
