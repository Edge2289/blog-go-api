package Common

import (
	"blog-go-api/app/config"
	imgsModel "blog-go-api/app/model/imgs"
	"blog-go-api/utils/json"
	uploadimg "blog-go-api/utils/storage/qiniu"
	"blog-go-api/utils/tools"

	//jsonc "encoding/json"
	"github.com/gin-gonic/gin"
	//"io/ioutil"
	"net/http"
	"strconv"
)

type Img struct {
	Ids string `json:"ids"`
	Data   string `json:"data"`
	CateId  int `json:"cate_id"`
}

/**
上传图片
*/
func UploadImg(c *gin.Context) {

	var img imgsModel.Imgs
	var imgData Img
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//jsonc.Unmarshal(data, &img)

	_ = c.BindJSON(&imgData)
	// 上传七牛云
	name, err := uploadimg.Upload([]byte(imgData.Data))
	utilGin := json.Gin{Ctx: c}
	if err != nil {
		utilGin.Fail(http.StatusBadRequest, "-11000", nil)
		return
	}
	img.Url = name
	img.CateId, _ = strconv.Atoi(c.Request.FormValue("cate_id"))
	img.Domain = config.ImgDomain

	i, err := img.AddImg()
	if err != nil {
		utilGin.Fail(http.StatusBadRequest, "-1102", nil)
		return
	}
	if i {
		utilGin.Success(http.StatusOK, "", nil)
		return
	}
	utilGin.Fail(http.StatusBadRequest, "-1103", nil)
}

/**
获取图片
*/
func GetImg(c *gin.Context) {

	var img imgsModel.Imgs
	img.CateId, _ = strconv.Atoi(c.Request.FormValue("cate_id"))
	page := tools.GetPage(c)
	pageSize := tools.GetPageSize(c)
	data, count, err := img.GetImgs(page, pageSize)
	utilGin := json.Gin{Ctx: c}
	if err != nil {
		utilGin.Fail(http.StatusBadRequest, "-1101", nil)
		return
	}
	utilGin.PageOK(data, count, page, pageSize)
}

/**
删除图片
*/
func DelImg(c *gin.Context) {

	var imgModel imgsModel.Imgs
	var img Img
	_ = c.BindJSON(&img)

	i, err := imgModel.DelImg(img.Ids)
	utilGin := json.Gin{Ctx: c}
	if err != nil {
		utilGin.Fail(http.StatusBadRequest, "-1102", nil)
		return
	}
	if i {
		utilGin.Success(http.StatusOK, "", nil)
		return
	}
	utilGin.Fail(http.StatusBadRequest, "-1102", nil)
}

/**
移动图片分组
*/
func MvImgCategory(c *gin.Context) {

	var imgModel imgsModel.Imgs
	var img Img
	_ = c.BindJSON(&img)
	imgModel.CateId = img.CateId
	i, err := imgModel.UpdateImsCategory(img.Ids)

	utilGin := json.Gin{Ctx: c}
	if err != nil {
		utilGin.Fail(http.StatusBadRequest, "-1102", nil)
		return
	}
	if i {
		utilGin.Success(http.StatusOK, "", nil)
		return
	}
	utilGin.Fail(http.StatusBadRequest, "-1102", nil)
}
