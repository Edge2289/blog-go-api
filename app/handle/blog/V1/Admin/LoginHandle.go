package Admin

import (
	"blog-go-api/app/middleware/jwt"
	"blog-go-api/app/model/admin"
	"blog-go-api/utils/json"
	"fmt"

	"blog-go-api/app/validators/adminValidators"
	"blog-go-api/utils/pkg"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

// 验证码程序
var store = base64Captcha.DefaultMemStore

/**
登陆处理层
*/
func Login(c *gin.Context) {
	var loginVals admin.LoginData
	//var jwtData = make(map[string]interface{})
	// 获取并且解析数据
	_ = c.BindJSON(&loginVals)
	fmt.Println("loginVals")
	fmt.Println(loginVals)

	if vErr := adminValidators.LoginValidators(loginVals); vErr != "" {
		pkg.AssertCode(http.StatusBadRequest, vErr, c)
	}
	//获取参数
	//loginVals.Username = c.Request.FormValue("user_name")
	//loginVals.Password = c.Request.FormValue("pwd")

	if loginVals.Code == "" {
		pkg.AssertCode(http.StatusBadRequest, "验证码不正确", c)
	}
	if loginVals.Username == "" || loginVals.Password == "" {
		pkg.AssertCode(http.StatusBadRequest, "-10001", c)
	}

	// 验证  验证码操作   错误码 -10002
	if !store.Verify(loginVals.UUID, loginVals.Code, true) {
		pkg.Assert(false, "-10002", c)
	}
	// 登陆操作
	adminData, err := loginVals.LoginGetUserList()
	pkg.AssertErr(err, "-10001", c)
	//// 生成token 生成token 失败
	token, err := jwt.GetJwtTokenMiddleware(adminData.Id)
	pkg.AssertErr(err, "-10003", c)

	utilGin := json.Gin{Ctx: c}
	// 设置token header 头
	c.Header("Access-Control-Expose-Headers", "Authorization")
	c.Header("Authorization", "Bearer "+token)

	utilGin.Success(http.StatusOK, "", adminData)
}
