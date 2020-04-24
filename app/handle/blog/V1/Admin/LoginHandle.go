package Admin

import (
	"blog-go-api/app/middleware/jwt"
	"blog-go-api/app/model/admin"
	"blog-go-api/app/validators/adminValidators"
	"blog-go-api/utils/json"
	"blog-go-api/utils/pkg"
	jsonc "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"io/ioutil"
	"net/http"
)

// 验证码程序
var store = base64Captcha.DefaultMemStore
/**
 登陆处理层
 */
func Login(c *gin.Context) {
	var loginVals admin.LoginData
	utilGin := json.Gin{Ctx: c}
	//var jwtData = make(map[string]interface{})
	// 获取并且解析数据
	data, _ := ioutil.ReadAll(c.Request.Body)
	jsonc.Unmarshal(data, &loginVals)

	if vErr := adminValidators.LoginValidators(loginVals); vErr != ""  {
		pkg.AssertCode(http.StatusBadRequest,vErr, c)
	}

	//获取参数
	//loginVals.Username = c.Request.FormValue("user_name")
	//loginVals.Password = c.Request.FormValue("pwd")

	if loginVals.Code == "" {
		pkg.AssertCode(http.StatusBadRequest,"验证码不正确", c)
	}
	if loginVals.Username == "" || loginVals.Password == "" {
		pkg.AssertCode(http.StatusBadRequest,"-10001", c)
	}

	fmt.Println(loginVals)
	/**
	 先验证  验证码是否正确
			再判断账号密码是否正确
	 */

	// 验证  验证码操作   错误码 -10002
	if !store.Verify(loginVals.UUID, loginVals.Code, true) {
		pkg.Assert(false,"-10002", c)
	}
	// 登陆操作
	adminData, err := loginVals.LoginGetUserList()
	pkg.AssertErr(err,"-10001", c)
	//jwtData["uid"] = adminData.Id
	//fmt.Println(jwtData)

	//var roleData []int
	//for _, v := range adminData.RoleData {
	//	roleData = append(roleData, v.RId)
	//}
	//jwtData["roleData"] = roleData
	//fmt.Println(jwtData)

	//// 生成token 生成token 失败
	token, err := jwt.GetJwtTokenMiddleware(adminData.Id)
	pkg.AssertErr(err,"-10003", c)
	fmt.Println(token)

	// 设置token header 头
	c.Header("Access-Control-Expose-Headers", "Authorization");
	c.Header("Authorization", "Bearer " + token)

	utilGin.Success(http.StatusOK, "", adminData)
}