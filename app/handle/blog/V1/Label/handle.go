package Label

import (
	labelModel "blog-go-api/app/model/label"
	labelService "blog-go-api/app/service/blog/v1/label"
	"blog-go-api/utils/json"
	"blog-go-api/utils/tools"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	label.IsState, _ = strconv.Atoi(c.Request.FormValue("is_state"))
	data, err := labelService.AddLabel(label)
	utilGin := json.Gin{Ctx: c}
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

	label.Id, _ = strconv.Atoi(c.Request.FormValue("id"))
	label.Label = c.Request.FormValue("label")
	label.Color = c.Request.FormValue("color")
	label.IsState, _ = strconv.Atoi(c.Request.FormValue("is_state"))
	fmt.Println("label", label)
	data, err := labelService.UpdateLabel(label)
	utilGin := json.Gin{Ctx: c}
	if err != nil || !data {
		utilGin.Fail(http.StatusBadRequest, "-1101", nil)
		return
	}
	utilGin.Success(http.StatusOK, "", data)
}

/**
删除
*/
func DelLabel(c *gin.Context) {

	var label labelModel.Label
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//jsonc.Unmarshal(data, &label)

	label.Id, _ = strconv.Atoi(c.Request.FormValue("id"))
	fmt.Println(label)
	data, err := labelService.DelLabel(label)
	utilGin := json.Gin{Ctx: c}
	if err != nil || !data {
		utilGin.Fail(http.StatusBadRequest, "-1101", nil)
		return
	}
	utilGin.Success(http.StatusOK, "", data)
}

/**
查询
*/
func GetLabel(c *gin.Context) {
	var label labelModel.Label
	//data, _ := ioutil.ReadAll(c.Request.Body)
	//jsonc.Unmarshal(data, &label)

	label.Id, _ = strconv.Atoi(c.Request.FormValue("id"))
	label.Label = c.Request.FormValue("label")
	label.IsState, _ = strconv.Atoi(c.Request.FormValue("is_state"))

	page := tools.GetPage(c)
	pageSize := tools.GetPageSize(c)

	data, count, err := labelService.GetLabel(label, page, pageSize)
	utilGin := json.Gin{Ctx: c}
	if err != nil {
		utilGin.Fail(http.StatusBadRequest, "-1101", data)
		return
	}
	utilGin.PageOK(data, count, page, pageSize)
}
