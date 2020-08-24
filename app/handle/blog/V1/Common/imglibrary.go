package Common

import (
	"blog-go-api/app/model/imgs"
	"blog-go-api/app/service/blog/v1/common"
	"blog-go-api/utils/json"
	//jsonc "encoding/json"
	"github.com/gin-gonic/gin"
	//"io/ioutil"
	"net/http"
	"strconv"
)

/**
获取图片分组
*/
func GetImgGroup(c *gin.Context) {
	/**
	获取前端数据
	*/
	data, err := common.GetImgGroup()

	utilGin := json.Gin{Ctx: c}
	if err != nil {
		utilGin.Fail(http.StatusBadRequest, "-1101", nil)
		return
	}
	utilGin.Success(http.StatusOK, "", data)
}

/**
删除图片分组
*/
func DelImgGroup(c *gin.Context) {

	var img imgs.ImgsCategory
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//jsonc.Unmarshal(data, &img)

	img.Id, _ = strconv.Atoi(c.Request.FormValue("id"))
	bool, _ := common.DelImgGroup(img)

	utilGin := json.Gin{Ctx: c}
	if !bool {
		utilGin.Fail(http.StatusBadRequest, "-1102", nil)
		return
	}
	utilGin.Success(http.StatusOK, "", "")
}

/**
更新图片分组
*/
func UpdateImgGroup(c *gin.Context) {

	var img imgs.ImgsCategory
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//jsonc.Unmarshal(data, &img)

	img.Id, _ = strconv.Atoi(c.Request.FormValue("id"))
	img.Name = c.Request.FormValue("name")
	img.Sort, _ = strconv.Atoi(c.Request.FormValue("sort"))

	bool, err := common.UpdateImgGroup(img)
	utilGin := json.Gin{Ctx: c}

	if err != nil {
		utilGin.Fail(http.StatusBadRequest, "-1102", nil)
		return
	}
	if bool {
		utilGin.Success(http.StatusOK, "", "")
		return
	}
	utilGin.Fail(http.StatusBadRequest, "-1102", nil)
}

/**
新增图片分组
*/
func AddImgGroup(c *gin.Context) {

	var img imgs.ImgsCategory
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//jsonc.Unmarshal(data, &img)
	img.Name = c.Request.FormValue("name")
	img.Sort, _ = strconv.Atoi(c.Request.FormValue("sort"))

	bool, err := common.AddImgGroup(img)

	utilGin := json.Gin{Ctx: c}
	if err != nil {
		utilGin.Fail(http.StatusBadRequest, "-1102", nil)
		return
	}
	if bool {
		utilGin.Success(http.StatusOK, "", "")
		return
	}
	utilGin.Fail(http.StatusBadRequest, "-1102", nil)
}
