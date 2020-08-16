package router

import (
	"blog-go-api/app/config"
	"blog-go-api/app/middleware"
	"blog-go-api/utils/json"
	"fmt"
	"log"
	"net/http"

	//"fmt"
	"github.com/gin-gonic/gin"
)

var Routers *gin.Engine
var apiRouter string

func Run() *gin.Engine {
	apiRouter = config.AppBlogRouter
	r := gin.Default()

	// 记录日志
	r.Use(middleware.LogsRequestToFile())
	// 设置header 头
	r.Use(Cors()) //开启中间件 允许使用跨域请求
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
	AdminRouter("admin")  // 管理员
	SystemRouter("system")
	ArticleRouter("article") // 文章

	// 开启端口
	r.Run(config.AppPort)
	return r
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}