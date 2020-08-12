package ArticleHandle

import (
	"blog-go-api/app/model/article"
	"blog-go-api/utils/json"
	"blog-go-api/utils/tools"
	//jsonc "encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * 获取文章
 */
func Get(c *gin.Context) {

	var article article.Article
	//article.CateId = 1
	//article.Title = 1
	//article.Nick = 1
	//article.IsState =

	// 获取并且解析数据
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//jsonc.Unmarshal(data, &article)

	article.Title = c.Query("title")
	article.CateId, _ = tools.StringToInt(c.Query("cate_id"))
	article.Nick = c.Query("nick")
	article.IsState, _ = tools.StringToInt(c.Query("is_state"))
	label := c.Query("label_id")

	page, _ := tools.StringToInt(c.Query("page"))
	pageSize, _ := tools.StringToInt(c.Query("page_size"))
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 30
	}
	reData, count, err := article.GetSearchList(label, page, pageSize)

	/**
	常規返回
	 */
	utilGin := json.Gin{Ctx: c}
	if err != nil {
		utilGin.Fail(http.StatusBadRequest, "-1101", nil)
		return
	}
	utilGin.PageOK(reData, count, page, pageSize)
}

