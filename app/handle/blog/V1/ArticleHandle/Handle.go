package ArticleHandle

import (
	"blog-go-api/app/model/article"
	"blog-go-api/utils/json"
	"blog-go-api/utils/tools"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var addArticleJson struct {
	Id           int        `json:"id"`
	Title        string     `gorm:"column:title;" json:"title"`
	Describe     string     `gorm:"column:describe;" json:"describe"`
	Img          string     `gorm:"column:img;" json:"img"`
	Nick         string     `gorm:"column:nick;" json:"nick"`
	CateId       int        `gorm:"column:cate_id;" json:"cate_id"`
	Introduction int     `gorm:"column:introduction;" json:"introduction"`
	IsComment    int        `gorm:"column:is_comment;" json:"is_comment"`
	IsState      int        `gorm:"column:is_state;" json:"is_state"`
	ClickNum     int        `gorm:"column:click_num;" json:"click_num"`
	ReadNum      int        `gorm:"column:read_num;" json:"read_num"`
	Text      string        `gorm:"column:text;" json:"text"`
	Markdown      string    `gorm:"column:markdown;" json:"markdown"`
	ShowType      int        `gorm:"column:is_state;" json:"show_type"`
	// 标签
	LabelData []int `json:"labels"`
}

/**
 * 获取文章
 */
func Get(c *gin.Context) {

	var article article.Article
	// 获取并且解析数据
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//jsonc.Unmarshal(data, &article)

	article.Title = c.Query("title")
	article.Id, _ = tools.StringToInt(c.Query("id"))
	article.CateId, _ = tools.StringToInt(c.Query("cate_id"))
	article.Nick = c.Query("nick")
	article.IsState, _ = tools.StringToInt(c.Query("is_state"))
	label := c.Query("label_id")

	page := tools.GetPage(c)
	pageSize := tools.GetPageSize(c)
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

/**
  post 请求
*/
func Post(c *gin.Context) {

	var addArticle = addArticleJson

	var articleDetails article.Article
	_ = c.BindJSON(&addArticle)

	articleDetails.Title = addArticle.Title
	articleDetails.Describe = addArticle.Describe
	articleDetails.Img = addArticle.Img
	articleDetails.Nick = addArticle.Nick
	articleDetails.CateId = addArticle.CateId
	articleDetails.IsComment =addArticle.IsComment
	articleDetails.IsState = addArticle.IsState
	articleDetails.ClickNum = addArticle.ClickNum
	articleDetails.ReadNum = addArticle.ReadNum

	articleDetails.TextData = article.ArticleText {
		Text: addArticle.Text,
		Markdown: addArticle.Markdown,
	}

	var labelData[] article.ArticleLabel

	for _, id := range addArticle.LabelData {
		labelData = append(labelData, article.ArticleLabel{
			LabelId: id,
			OperatorId: 1,
			OperatorName: "系统",
			CreateTime:   time.Now(),
		})
	}
	articleDetails.LabelData = labelData
	state, err := articleDetails.AddArticleDetail()
	/**
	常規返回
	*/
	utilGin := json.Gin{Ctx: c}
	if err != nil || state != true {
		utilGin.Fail(http.StatusBadRequest, "-1101", nil)
		return
	}
	utilGin.Success(http.StatusOK, "", nil)
}
/**
更新
*/
func Put(c *gin.Context) {

	var updateArticle = addArticleJson

	var articleDetails article.Article
	_ = c.BindJSON(&updateArticle)

	articleDetails.Id = updateArticle.Id
	articleDetails.Title = updateArticle.Title
	articleDetails.Describe = updateArticle.Describe
	articleDetails.Img = updateArticle.Img
	articleDetails.Nick = updateArticle.Nick
	articleDetails.CateId = updateArticle.CateId
	articleDetails.IsComment =updateArticle.IsComment
	articleDetails.IsState = updateArticle.IsState
	articleDetails.ClickNum = updateArticle.ClickNum
	articleDetails.ReadNum = updateArticle.ReadNum

	articleDetails.TextData = article.ArticleText {
		Text: updateArticle.Text,
		Markdown: updateArticle.Markdown,
	}

	var labelData[] article.ArticleLabel

	for _, id := range updateArticle.LabelData {
		labelData = append(labelData, article.ArticleLabel{
			ArticleId: articleDetails.Id,
			LabelId: id,
			OperatorId: 1,
			OperatorName: "系统",
			CreateTime:   time.Now(),
		})
	}
	articleDetails.LabelData = labelData
	state, err := articleDetails.UpdateArticleDetail()

	utilGin := json.Gin{Ctx: c}
	if err != nil || state != true {
		utilGin.Fail(http.StatusBadRequest, "-1101", nil)
		return
	}
	utilGin.Success(http.StatusOK, "", nil)
}

/**
删除
*/
func Delete(c *gin.Context) {

	var articleDetails article.Article
	_ = c.BindJSON(&articleDetails)
	state, err := articleDetails.DelArticleDetail()

	utilGin := json.Gin{Ctx: c}
	if err != nil || state != true {
		utilGin.Fail(http.StatusBadRequest, "-1101", nil)
		return
	}
	utilGin.Success(http.StatusOK, "", nil)
}
