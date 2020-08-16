package tools

import "github.com/gin-gonic/gin"

/**
 獲取頁碼
 */
func GetPage(c *gin.Context) int {

	page, _ := StringToInt(c.Query("page"))
	if page == 0 {
		page = 1
	}
	return page
}
/**
	獲取頁數
 */
func GetPageSize(c *gin.Context) int {
	pageSize, _ := StringToInt(c.Query("page_size"))
	if pageSize == 0 {
		pageSize = 30
	}
	return pageSize
}