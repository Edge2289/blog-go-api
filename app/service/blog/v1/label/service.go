package label

import (
	labelModel "blog-go-api/app/model/label"
	"github.com/gin-gonic/gin"
	"strconv"
)

/**
	新增
*/
func AddLabel(label labelModel.Label) (bool, error) {

	i ,error := label.AddLabel()
	return i, error
}

/**
	更新
*/
func UpdateLabel(c *gin.Context) {

}

/**
	删除
*/
func DelLabel(c *gin.Context) {

}

/**
	查询
*/
func GetLabel(c *gin.Context) {

}
