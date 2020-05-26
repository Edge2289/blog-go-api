package Label

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	labelService "blog-go-api/app/service/blog/v1/label"
)

/**
新增
*/
func AddLabel(c *gin.Context) {

	var label labelModel.Label
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//jsonc.Unmarshal(data, &label)

	label.Label = c.Request.FormValue("label")
	label.Color = c.Request.FormValue("color")
	label.IsState,_ = strconv.Atoi(c.Request.FormValue("is_state"))
	data, err := labelService.AddLabel(label);
	if err != nil || !data {
		utilGin.Fail(http.StatusBadRequest, "-1101", nil)
		return
	}
	utilGin.Success(http.StatusOK, "", data)
}

/**
更新
 */
func UpdateLabel(c *gin.Context) {

	var label labelModel.Label
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//jsonc.Unmarshal(data, &label)

	label.Id,_ = strconv.Atoi(c.Request.FormValue("id"))
	label.Label = c.Request.FormValue("label")
	label.Color = c.Request.FormValue("color")
	label.IsState,_ = strconv.Atoi(c.Request.FormValue("is_state"))

}

/**
删除
 */
func DelLabel(c *gin.Context) {

	var label labelModel.Label
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//jsonc.Unmarshal(data, &label)

	label.Id,_ = strconv.Atoi(c.Request.FormValue("id"))

}

/**
查询
 */
func GetLabel(c *gin.Context) {
	var label labelModel.Label
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//jsonc.Unmarshal(data, &label)

	label.Label = c.Request.FormValue("label")
	label.IsState,_ = strconv.Atoi(c.Request.FormValue("is_state"))

}