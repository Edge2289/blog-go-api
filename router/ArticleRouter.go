package router

import (
	"blog-go-api/app/handle/blog/V1/ArticleHandle"
	"blog-go-api/app/handle/blog/V1/CateHandle"
	"blog-go-api/app/handle/blog/V1/Label"
	//"blog-go-api/app/middleware"
)

/**
  文章的全部路由
		文章   /
		文章分类  cate
		文章标签  label
*/
func ArticleRouter(base string) {

	r := Routers.Group(apiRouter + "v1/" + base)

	/**
		检查请求权限   以及JWT
		下次JWT 和  检查Auth 拆分出来
		前端请求数据加密对比

	middleware.CheckJwt()
	*/
	r.Use()
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
