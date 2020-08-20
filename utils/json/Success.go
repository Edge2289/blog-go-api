package json

import (
	"blog-go-api/resources/lang"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	Ctx *gin.Context
}

/**
 * 定义返回类型
 */
type ResponseData struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

/**
 * 定义返回类型
 */
type PageResponseData struct {
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
	Count    int         `json:"count"`
	Data     interface{} `json:"list"`
}

/**
 * 成功返回值
 */
func (g *Gin) Success(code int, msg string, data interface{}) {
	if msg == "" {
		msg = "操作成功"
	}
	g.Ctx.JSON(200, ResponseData{
		Code:    code,
		Message: lang.GetLang(msg),
		Data:    data,
	})
	return
}

/**
 * 失败返回值
 */
func (g *Gin) Fail(code int, msg string, data interface{}) {
	if msg == "" {
		msg = "操作失败"
	}
	g.Ctx.JSON(200, ResponseData{
		Code:    code,
		Message: lang.GetLang(msg),
		Data:    data,
	})
	return
}

// 分页数据处理
func (g *Gin) PageOK(result interface{}, count int, pageIndex int, pageSize int) {
	var res PageResponseData
	res.Data = result
	res.Count = count
	res.Page = pageIndex
	res.PageSize = pageSize

	g.Ctx.JSON(200, ResponseData{
		Code:    200,
		Message: "操作成功",
		Data:    res,
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
