package CateHandle

import (
	"blog-go-api/app/model/category"
	"blog-go-api/utils/json"
	"blog-go-api/utils/time"
	"blog-go-api/utils/tools"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get(c *gin.Context) {

	var categoryModel  category.Category

	categoryModel.Id, _ = tools.StringToInt(c.Query("id"))
	categoryModel.Name = c.Query("name")
	categoryModel.IsState, _ = tools.StringToInt(c.Query("is_state"))

	page := tools.GetPage(c)
	pageSize := tools.GetPageSize(c)
	reData, count, err := categoryModel.GetList(page, pageSize)
	/**
	常規返回
	*/
	utilGin := json.Gin{Ctx: c}
	if err != nil {
		utilGin.Fail(http.StatusBadRequest, "-1101", nil)
		return
	}
	utilGin.PageOK(reData, count, page, pageSize)
}
/**
新增分類
*/
func AddCategory(c *gin.Context)  {
	var categoryModel  category.Category

	c.BindJSON(&categoryModel)
	categoryModel.OperatorId = 1
	categoryModel.OperatorName = "ces"
	categoryModel.CreateTime = time.GetDatabaseDate()
	state, err := categoryModel.AddCategory()
	/**
	常規返回
	*/
	utilGin := json.Gin{Ctx: c}
	if err != nil || state != true{
		utilGin.Fail(http.StatusBadRequest, "-1101", nil)
		return
	}
	utilGin.Success(http.StatusOK, "", nil)
}

/**
	修改分類
*/
func UpdateCategory(c *gin.Context)  {
	var categoryModel  category.Category

	c.BindJSON(&categoryModel)
	categoryModel.UpdateTime = time.GetDatabaseDate()
	state, err := categoryModel.UpdateCategory()
	///**
	//常規返回
	//*/
	utilGin := json.Gin{Ctx: c}
	if err != nil || state != true{
		utilGin.Fail(http.StatusBadRequest, "-1101", nil)
		return
	}
	utilGin.Success(http.StatusOK, "", nil)
}

func DelCategory(c *gin.Context)  {

	var categoryModel  category.Category

	utilGin := json.Gin{Ctx: c}
	c.BindJSON(&categoryModel)
	if categoryModel.Id == 0 {
		utilGin.Fail(http.StatusBadRequest, "-12000", nil)
	}

	state, err := categoryModel.DelCategory()
	/**
	常規返回
	*/
	if err != nil || state != true{
		utilGin.Fail(http.StatusBadRequest, "-1101", nil)
		return
	}
	utilGin.Success(http.StatusOK, "", nil)
}