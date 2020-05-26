package imgutils

import (
	imgsModel "blog-go-api/app/model/imgs"
	"blog-go-api/utils/json"
	"fmt"
	"testing"
)

// 查询图片分类
func TestGetArtcle(t *testing.T) {
	var img imgsModel.ImgsCategory
	data, err := img.GetImgsCategorys()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("data", data)
}

// 新增图片分类
func TestAddArticle(t *testing.T) {

	var addImgs imgsModel.ImgsCategory
	addImgs.Name = "两极反转啊"
	addImgs.Sort = 70
	inData, _ := addImgs.AddImsCategory()
	fmt.Println(inData)
}

// 更新图片分类
func TestUpdateArticle(t *testing.T) {

	var updateImgs imgsModel.ImgsCategory
	updateImgs.Id = 2
	updateImgs.Name = "测试分类"
	datas, _ := updateImgs.UpdateImsCategory()
	fmt.Println("datas", datas)
}

// 软删除图片分类
func TestDelArticle(t *testing.T)  {
	var delImgs imgsModel.ImgsCategory
	delImgs.Id = 4
	delData, _ := delImgs.DelImsCategory()
	fmt.Println(delData)
}