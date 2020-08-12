package article

import (
	"blog-go-api/app/model/article"
	"fmt"
	"testing"
	"time"
)

type ArticleLabelTest struct {
	ArticleId int
	LabelId int
}

/**
	获取搜索文章列表
 */
func TestArticleList(t *testing.T)  {

	var article article.Article
	//article.Title = "測試"
	data, _ , _ := article.GetSearchList("1", 1, 30)
	for _, v := range data  {
		fmt.Println("-------------------")
		fmt.Println(v)
		fmt.Println("-------------------")
	}
	fmt.Println(data)
}

/**
	测试获取文档
 */
func TestArticleDetail(t *testing.T) {
	var article article.Article
	article.Id = 1
	data, _ := article.GetArticleDetail()
	fmt.Println(data)
}

/**
	新增文档
 */
func TestAddArticleDetail(t *testing.T) {

	var articleText article.ArticleText
	var articleBi article.Article
	//var articleLabel article.ArticleLabel

	articleBi.Title = "测试新增文档TitleaaaASasa"
	articleBi.Describe = "测试新增文档DescriASasbe1123"
	articleBi.Img = "测试新增文档ImgasS12123"
	articleBi.Nick = "测试新增文档aSasNick12313"
	articleBi.CateId = 1
	articleBi.Introduction = "测试新SsSas增文档12313Introduction"
	articleBi.IsComment = 1
	articleBi.IsState = 1
	articleBi.ClickNum = 1
	articleBi.ReadNum = 1

	// 明细
	articleText.Text = "Textxxxx"
	articleText.Markdown = "Markdownasdasda"
	
	// 标签

	label := []article.ArticleLabel{
		article.ArticleLabel{
			6,1,1,1,"1",time.Now(),time.Now(),nil,
		},
		article.ArticleLabel{
			7,1,2,1,"1",time.Now(),time.Now(),nil,
		},
	}

	articleBi.TextData = articleText
	articleBi.LabelData = label

	bool, err := articleBi.AddArticleDetail()
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}
	fmt.Println(bool)
}