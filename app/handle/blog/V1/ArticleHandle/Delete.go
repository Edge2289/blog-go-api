package ArticleHandle

import "github.com/gin-gonic/gin"

/**
删除
*/
func Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "ArticleHandle",
	})
}
