package middleware

import "github.com/gin-gonic/gin"

/**
 * 根据用户检查是否有权限操作
 * 权限检查中间件
 */
func CheckAuthRole() gin.HandlerFunc  {
	 return func(c *gin.Context) {
		 //根据上下文获取载荷claims 从claims获得role
		 //claims := c.MustGet("claims").(*jwt.MapClaims)
	 }
}