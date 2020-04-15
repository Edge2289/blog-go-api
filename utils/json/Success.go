package json

import (
	"github.com/gin-gonic/gin"
	"encoding/json"
	"blog-go-api/resources/lang"
)

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
	if msg == ""{
		msg = "操作成功。"
	}
	g.Ctx.JSON(200, ResponseData{
		Code    : code,
		Message : lang.GetLang(msg),
		Data    : data,
	})
	return
}


/**
 * 失败返回值
 */
func (g *Gin) Fail(code int, msg string, data interface{}) {
	if msg == ""{
		msg = "操作失败。"
	}
	g.Ctx.JSON(400, ResponseData{
		Code    : code,
		Message : lang.GetLang(msg),
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