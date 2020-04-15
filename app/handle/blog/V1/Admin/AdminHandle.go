package Admin

import (
	AdminModel "blog-go-api/app/model/admin"
	"blog-go-api/utils/json"
	//"blog-go-api/utils/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 新增管理员
 */
func AddAdmin(c *gin.Context)  {
	var admin AdminModel.Admin
	//err := c.ShouldBind(&admin)
	//pkg.AssertErr(err,"", c)
	//if admin.LoginName == "" {
	//	pkg.Assert(false,"抱歉未找到相关信息", c)
	//}
	fmt.Println(c.Request.FormValue("login_name"))
	admin.LoginName = c.Request.FormValue("login_name")
	admin.Password = c.Request.FormValue("pwd")
	fmt.Println(admin)
	utilGin := json.Gin{Ctx: c}
	utilGin.Success(http.StatusOK, "-1000", nil)
	fmt.Println(123123)
}