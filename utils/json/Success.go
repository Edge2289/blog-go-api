package json

import "github.com/gin-gonic/gin"
import "encoding/json"

type Gin struct {
	Ctx *gin.Context
}
/**
 * 定义返回类型
 */
type ResponseData struct {
	Code     int         `json:"code"`
	Message  string      `json:"msg"`
	Data     interface{} `json:"data"`
}

/**
 * 成功返回值
 */
func (g *Gin) Success(code int, msg string, data interface{}) {
	g.Ctx.JSON(200, ResponseData{
		Code    : code,
		Message : msg,
		Data    : data,
	})
	return
}


/**
 * 失败返回值
 */
func (g *Gin) Fail(code int, msg string, data interface{}) {
	g.Ctx.JSON(400, ResponseData{
		Code    : code,
		Message : msg,
		Data    : data,
	})
	return
}


func Encode(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}