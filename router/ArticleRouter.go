package router

import (
	"blog-go-api/app/handle/blog/V1/ArticleHandle"
	"blog-go-api/app/handle/blog/V1/CateHandle"
	"blog-go-api/app/handle/blog/V1/Label"
	jwtAuth "blog-go-api/app/middleware"

	//"blog-go-api/app/middleware"
)

/**
  文章的全部路由
		文章   /
		文章分类  cate
		文章标签  label
*/
func ArticleRouter(base string) {

	/**
	需要鉴权的
	*/
	RoutersAuth := Routers.Group(apiRouter + "v1")
	RoutersAuth.Use(jwtAuth.CheckJwt())

	r := RoutersAuth.Group("/" +base)
	{
		// 文章
		r.GET("/list", ArticleHandle.Get) // 获取文章
		r.POST("/add", ArticleHandle.Post)
		r.DELETE("del", ArticleHandle.Delete)
		r.PUT("/put", ArticleHandle.Put)

		// 文章分类
		r.GET("/cate", CateHandle.Get) // 获取文章
		r.POST("/cate", CateHandle.AddCategory)
		r.DELETE("/cate", CateHandle.DelCategory)
		r.PUT("/cate", CateHandle.UpdateCategory)

		// 文章标签
		r.GET("/label", Label.GetLabel) // 获取文章
		r.POST("/label", Label.AddLabel)
		r.DELETE("/label", Label.DelLabel)
		r.PUT("/label", Label.UpdateLabel)
	}
}
