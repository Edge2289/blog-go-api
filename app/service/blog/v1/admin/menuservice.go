package admin

import (
	"blog-go-api/app/model/admin"
	jsonc "encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

/**
获取菜单数据
*/
func GetMenuList(c *gin.Context) ([]admin.MenuModel, string) {
	var menuData admin.MenuModel
	/**
	获取传过来的数据
	 验证一下
	*/
	// 获取并且解析数据
	data, _ := ioutil.ReadAll(c.Request.Body)
	jsonc.Unmarshal(data, &menuData)

	data, err := menuData.GetMenuList(1, 10)
	if err != nil {
		return nil, "-1101"
	}
	return data, ""
}

func AddMenuList() {

	//adminValidators.menuValidator()
}
