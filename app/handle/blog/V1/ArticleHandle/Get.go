package ArticleHandle

import "github.com/gin-gonic/gin"

/**
 * 获取文章
 */
func Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg" : "ArticleHandle",
	})
}

