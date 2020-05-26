package middleware

import (
	"blog-go-api/app/middleware/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"

	"blog-go-api/app/model/admin"
	"blog-go-api/utils/pkg"
	//"github.com/gin-gonic/gin"
)

/**
 * 根据用户检查是否有权限操作
 * 权限检查中间件
 */
func CheckJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		/**
		 先获取token
		 解析token  验证token
			1：过期直接返回
			2：没过期，没权限  返回
		    3：next()
		*/
		authToken := c.Request.Header.Get("Authorization")
		if authToken == "" {
			pkg.AssertCode(http.StatusBadRequest, "-1003", c)
		}
		userData, err := jwt.JwtParseUser(c)
		pkg.AssertErr(err, "-1003", c)

		// userData  拿出role 去验证是否有该接口权限
		auth := checkUserAuth(userData.(float64), c.Request.URL.Path, c.Request.Method)
		pkg.Assert(auth, "-1003", c)
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		// 将当前请求的username信息保存到请求的上下文c上
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

/**
	检查是否有权限访问接口
 */
func checkUserAuth(userData float64, apiPath string, apiMethod string) bool {
	/**
		提取jwt 里面的用户信息
		有用户id  以及 用户角色
	 */
	// 先读redis里面的用户角色菜单

	// 没有则直接生成数据库的角色菜单
	adminMenuData, err := admin.GetAdminJwtMenu(int(userData))
	if err != nil{
		return false
	}
	for _, v := range  adminMenuData {
		if strings.ToUpper(v.ApiPath) == strings.ToUpper(apiPath) && strings.ToUpper(v.Method) == strings.ToUpper(apiMethod) {
			return true
		}
	}
	return false
}