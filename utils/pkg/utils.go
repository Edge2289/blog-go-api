package pkg

import (
	"blog-go-api/utils/json"
	"github.com/gin-gonic/gin"
	"strconv"
	//"strconv"
)

// Assert 条件断言
// 当断言条件为 假 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
func Assert(condition bool, msg string, c *gin.Context) {
	if !condition {
		statusCode := 400
		utilGin := json.Gin{Ctx: c}
		utilGin.Success(statusCode, msg, nil)
		c.Abort()
		panic("CustomErroe#" + strconv.Itoa(statusCode) + "#" + msg)
		return
	}
}

// AssertErr 错误断言
// 当 error 不为 nil 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
// 若 msg 为空，则默认为 error 中的内容
func AssertErr(err error, msg string, c *gin.Context) {
	if err != nil {
		statusCode := 400
		if msg == "" {
			msg = err.Error()
		}
		if msg == "-1003" {
			// 重新登录
			statusCode = 401
		}
		utilGin := json.Gin{Ctx: c}
		utilGin.Success(statusCode, msg, nil)
		c.Abort()
		panic("CustomErroe#" + strconv.Itoa(statusCode) + "#" + msg)
		return
	}
}

// Assert 条件断言
// 当断言条件为 假 时触发 panic
// 对于当前请求不会再执行接下来的代码，并且返回指定格式的错误信息和错误码
func AssertCode(code int, msg string, c *gin.Context) {
	if code != 200 {
		statusCode := code
		utilGin := json.Gin{Ctx: c}
		utilGin.Success(statusCode, msg, nil)
		c.Abort()
		panic("CustomErroe#" + strconv.Itoa(statusCode) + "#" + msg)
		return
	}
}
