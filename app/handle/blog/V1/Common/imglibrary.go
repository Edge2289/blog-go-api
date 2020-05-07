package Common

import (
	"blog-go-api/app/service/blog/v1/common"
	"blog-go-api/utils/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type imgs struct {

}

/**
	上传图片
 */
func UploadImg() {

}

/**
	删除图片
 */
func DelImgs() {

}

/**
获取图片分组
*/
func GetImgGroup(c *gin.Context) {
	/**
		获取前端数据
	 */
	utilGin := json.Gin{Ctx: c}
	common.GetImgGroup(c)
	var img = imgs.Imgs
	data, err := img.GetImgsCategorys()

	if err {
		utilGin.Fail(http.StatusBadRequest, "-1101", nil)
	}
	utilGin.StatusOK(http.StatusBadRequest, "", data)
}
/**
删除图片分组
*/
func DelImgGroup() {

	var img = imgs.Imgs
	data, err := img.GetImgsCategorys()

	if err {
		utilGin.Fail(http.StatusBadRequest, "-1101", nil)
	}
	utilGin.StatusOK(http.StatusBadRequest, "", data)
}
/**
更新图片分组
*/
func UpdateImgGroup() {

	var img = imgs.Imgs
	data, err := img.GetImgsCategorys()

	if err {
		utilGin.Fail(http.StatusBadRequest, "-1101", nil)
	}
	utilGin.StatusOK(http.StatusBadRequest, "", data)
}
/**
新增图片分组
*/
func AddImgGroup() {

	var img = imgs.Imgs
	data, err := img.GetImgsCategorys()

	if err {
		utilGin.Fail(http.StatusBadRequest, "-1101", nil)
	}
	utilGin.StatusOK(http.StatusBadRequest, "", data)
}