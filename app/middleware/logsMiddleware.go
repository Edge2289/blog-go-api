package middleware

import (
	"blog-go-api/app/config"
	jsonRequest "blog-go-api/utils/json"
	log "blog-go-api/utils/logger"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

var accessChannel = make(chan map[string]interface{})
var accessErrChannel = make(chan int)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

/**
记录每一次请求的日志
请求参数以及返回结果
*/
func LogsRequestToFile() gin.HandlerFunc {

	go handleAccessChannel()

	return func(c *gin.Context) {

		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		// 开始时间
		startTime := time.Now()

		// 请求参数
		requestData := c.Request.Form.Encode() // "请求参数"

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 返回参数
		reponseData := make(map[string]interface{})

		responseBody := bodyLogWriter.body.String()
		res := jsonRequest.ResponseData{}
		err := json.Unmarshal([]byte(responseBody), &res)
		if err == nil {
			reponseData["code"] = res.Code
			reponseData["Message"] = res.Message
			reponseData["Data"] = res.Data
		}

		//日志格式
		accessLogMap := make(map[string]interface{})

		accessLogMap["latencyTime"] = latencyTime
		accessLogMap["endTime"] = endTime
		accessLogMap["reqMethod"] = reqMethod
		accessLogMap["requestData"] = requestData
		accessLogMap["reponseData"] = responseBody
		accessLogMap["reqUri"] = reqUri
		accessLogMap["statusCode"] = statusCode
		accessLogMap["clientIP"] = clientIP

		//accessLogJson, _ := json.Marshal(accessLogMap)

		accessChannel <- accessLogMap
		// 状态码不等于正常 200
		//if statusCode != 200{
		//	accessErrChannel <- statusCode
		//}
	}
}

func handleAccessChannel() {

	for logData := range accessChannel {
		log.Addlog(logData, config.LogDirName+"request")
	}

	/**
	* 返回状态码不一样的时候执行
	 */
	//for statusCode := range accessErrChannel  {
	//	fmt.Println("error Code :", statusCode)
	//pushEmail(statusCode)
	//}
}

/**
推送邮箱
*/
func pushEmail(i int) {

}
