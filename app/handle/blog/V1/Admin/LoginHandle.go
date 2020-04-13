package Admin

import (
	"blog-go-api/app/model/admin"
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
 登陆处理层
 */
func Login(c *gin.Context)  {
	var admin admin.AdminModel
	admin.LoginName = c.Request.FormValue("user")
	admin.Password = c.Request.FormValue("pwd")
	verifCode := c.Request.FormValue("verifcode")
	fmt.Print(verifCode)
}