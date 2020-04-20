package Common

import "github.com/gin-gonic/gin"

/**
  添加访问日志
 */
func AddAccess(c *gin.Context)  {
	// 访问日志推送到es
	c.JSON(200, gin.H{
		"msg" : "ArticleHandle",
	})
}