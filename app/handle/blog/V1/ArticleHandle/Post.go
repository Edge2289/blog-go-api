package ArticleHandle

import "github.com/gin-gonic/gin"

/**
  post 请求
*/
func Post(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "ArticleHandle",
	})
}
