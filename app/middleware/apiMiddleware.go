package middleware

import (
	//jwtAuth "blog-go-api/app/middleware/jwt"
	//"github.com/dgrijalva/jwt-go"
	"blog-go-api/app/handle/blog/V1/Admin"
	//"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


/**
 登陆中间件
 1 判断用户的账号密码是否正确
		返回用户id 角色
 2 根据用户id 角色
		生成Token
 3 根据角色id 去查找菜单
		返回组装结构化 菜单
 4 返回前端用户以及菜单
 */
func LoginMiddleWare(c *gin.Context){
	Admin.Login(c)
}

/**
 退出中间件
 */
func LogoutMiddleWare()  {
	//var c *gin.Context
	//Admin.Logout
}

/**
 检查api是否正常请求
 */
func CheckApiToken() {

}
