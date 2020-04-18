package router

import (
	"blog-go-api/app/middleware"
	"blog-go-api/app/handle/blog/V1/ArticleHandle"
)

/**
  文章的全部路由
		文章   /
		文章分类  cate
		文章标签  label
 */
func ArticleRouter(base string)  {

	r := Routers.Group(apiRouter + "v1/" + base)

	/**
		检查请求权限   以及JWT
		下次JWT 和  检查Auth 拆分出来
		前端请求数据加密对比
	 */
	r.Use(middleware.CheckJwt())
	{
		// 文章
		r.GET("/", ArticleHandle.Get) // 获取文章
		r.POST("/", ArticleHandle.Post)
		r.DELETE("/", ArticleHandle.Delete)
		r.PUT("/", ArticleHandle.Put)

		// 文章分类
		r.GET("/cate", ArticleHandle.Get) // 获取文章
		r.POST("/cate", ArticleHandle.Get)
		r.DELETE("/cate", ArticleHandle.Get)
		r.PUT("/cate", ArticleHandle.Get)

		// 文章标签
		r.GET("/label", ArticleHandle.Get) // 获取文章
		r.POST("/label", ArticleHandle.Get)
		r.DELETE("/label", ArticleHandle.Get)
		r.PUT("/label", ArticleHandle.Get)
	}
}
