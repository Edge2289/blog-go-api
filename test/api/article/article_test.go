package article

import (
	"blog-go-api/app/model/article"
	"blog-go-api/utils/tools"
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

type ArticleLabelTest struct {
	ArticleId int
	LabelId   int
}

/**
获取搜索文章列表
*/
func TestArticleList(t *testing.T) {

	var article article.Article
	//article.Title = "測試"
	data, _, _ := article.GetSearchList("1", 1, 30)
	for _, v := range data {
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
	var workResultLock sync.WaitGroup
	for ii := 0; ii < 1000; ii++ {
		workResultLock.Add(1)
		go addArticle()
		fmt.Print("協程 開啓")
	}
	//主线程等待
	workResultLock.Wait()
	fmt.Print("結束")
}

func addArticle() {

	var articleText article.ArticleText
	var articleBi article.Article
	//var articleLabel article.ArticleLabel

	for i := 0; i < 50000; i++ {

		articleBi.Title = "测试新增文档TitleaaaASasa" + tools.IntToString(rand.Intn(1000000000000))
		articleBi.Describe = "测试新增文档DescriASasbe1123" + tools.IntToString(rand.Intn(1000000000000))
		articleBi.Img = "测试新增文档ImgasS12123" + tools.IntToString(rand.Intn(1000000000000))
		articleBi.Nick = "测试新增文档aSasNick12313" + tools.IntToString(rand.Intn(1000000000000))
		articleBi.CateId = rand.Intn(1000000)
		articleBi.Introduction = "测试新SsSas增文档12313Introduction" + tools.IntToString(rand.Intn(1000000000000))
		articleBi.IsComment = 1
		articleBi.IsState = 1
		articleBi.ClickNum = 1
		articleBi.ReadNum = 1

		// 明细
		articleText.Text = "Textxxxx" + tools.IntToString(rand.Intn(1000000000000)) + tools.IntToString(rand.Intn(1000000000000)) + tools.IntToString(rand.Intn(1000000000000)) + tools.IntToString(rand.Intn(1000000000000))
		articleText.Markdown = "Markdownasdasda" + tools.IntToString(rand.Intn(1000000000000)) + tools.IntToString(rand.Intn(1000000000000)) + tools.IntToString(rand.Intn(1000000000000))

		// 标签

		label := []article.ArticleLabel{
			article.ArticleLabel{
				ArticleId:    rand.Intn(100000),
				LabelId:      rand.Intn(100000),
				OperatorId:   rand.Intn(100000),
				OperatorName: "" + tools.IntToString(rand.Intn(100000)),
				CreateTime:   time.Now(),
			},
		}

		articleBi.TextData = articleText
		articleBi.LabelData = label

		s, _ := articleBi.AddArticleDetail()
		fmt.Print(s)
	}
}

/**

 */
func TestAddLabel(t *testing.T) {

	var workResultLock sync.WaitGroup
	for ii := 0; ii < 1000; ii++ {
		workResultLock.Add(1)
		//go addLabel()
		fmt.Print("協程 開啓")
	}
	//主线程等待
	workResultLock.Wait()
	fmt.Print("結束")
}
func addLabel() {
	//for i := 0; i < 10000;  i ++{
	//	var label label.Label
	//	label.Label =tools.IntToString(rand.Intn(15012)) + tools.IntToString(i)
	//	label.IsState = 1
	//	label.Color = "白" + tools.IntToString(i)
	//	label.OperatorId = rand.Intn(100000)
	//	label.OperatorName = ""+ tools.IntToString(rand.Intn(100000))
	//	label.CreateTime = time.Now()
	//	label.AddLabel()
	//}
}
