package ArticleHandle

import "github.com/gin-gonic/gin"

/**
更新
 */
func Put(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg" : "ArticleHandle",
	})
}