package Common

import (
	"blog-go-api/app/config"
	imgsModel "blog-go-api/app/model/imgs"
	"blog-go-api/utils/json"
	uploadimg "blog-go-api/utils/storage/qiniu"

	//jsonc "encoding/json"
	"github.com/gin-gonic/gin"
	//"io/ioutil"
	"net/http"
	"strconv"
)


/**
上传图片
*/
func UploadImg(c *gin.Context) {

	var img imgsModel.Imgs
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//jsonc.Unmarshal(data, &img)

	data := c.Request.FormValue("data")
	var imgData = []byte(data)
	// 上传七牛云
	name, err := uploadimg.Upload(imgData)
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
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//jsonc.Unmarshal(data, &img)

	img.CateId, _ = strconv.Atoi(c.Request.FormValue("cate_id"))
	page, _ := strconv.Atoi(c.Request.FormValue("page"))
	pageSize, _ := strconv.Atoi(c.Request.FormValue("pageSize"))
	data, err := img.GetImgs(page, pageSize)
	utilGin := json.Gin{Ctx: c}
	if err != nil {
		utilGin.Fail(http.StatusBadRequest, "-1101", nil)
		return
	}
	utilGin.Success(http.StatusOK, "", data)
}
/**
删除图片
*/
func DelImg(c *gin.Context) {

	var img imgsModel.Imgs
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//jsonc.Unmarshal(data, &img)

	img.CateId, _ = strconv.Atoi(c.Request.FormValue("cate_id"))
	i, err := img.DelImg()
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

	var img imgsModel.Imgs
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//jsonc.Unmarshal(data, &img)

	img.Id, _ = strconv.Atoi(c.Request.FormValue("id"))
	img.CateId, _ = strconv.Atoi(c.Request.FormValue("cate_id"))
	i, err := img.UpdateImsCategory()

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
