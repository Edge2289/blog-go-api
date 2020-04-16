package router

import (
	"blog-go-api/app/config"
	"blog-go-api/app/middleware"
	"blog-go-api/utils/json"
	"fmt"

	//"fmt"
	"github.com/gin-gonic/gin"
)

var Routers *gin.Engine

func Run() *gin.Engine {
	r := gin.Default()

	// 记录日志
	r.Use(middleware.LogsRequestToFile())
	// 设置header 头
	r.Use(middleware.NoCache)
	r.Use(middleware.Options)
	r.Use(middleware.Secure)


	fmt.Println("请求")
	// 一些不需要登录 ， 鉴权的接口
	// 这个是ping 返回
	r.GET("/ping", func(context *gin.Context) {
		//c.JSON(200, gin.H{
		//	"message": "pong",
		//})
		//context.Writer.WriteString("ping")
		utilGin := json.Gin{Ctx: context}
		var reData = map[string]interface{}{
			"user": "谢谢",
			"pwd":  "123456789",
			"menu": map[string]interface{}{
				"one" : 123123,
				"two" : "123123",
			},
		}
		utilGin.Success(200, "成功asdad", reData)
	})

	Routers = r

	// the jwt middleware
	//authMiddleware, _ := middleware.AuthInit()
	//fmt.Print(authMiddleware)
	CommonRouter()
	AdminRouter("admin")
	ArticleRouter("article") // 文章

	// 开启端口
	r.Run(config.AppPort)
	return r
}
