package logger

import (
	"blog-go-api/app/config"
	"errors"
	"reflect"
	"strings"

	logDrive "blog-go-api/utils/logger/drive"
	"fmt"
	//"strings"
)

/**
 填写日志的入口
 */
func Addlog(m map[string]interface{}, _dir string) {
	// 获取推送日志的驱动
	dirveConfig := config.LogDrive
	dirveConfigArr := strings.Split(dirveConfig,"/")

	dirveInterface := map[string]interface{}{
		"file" : logDrive.File,
		"es" : logDrive.Es,
		"mq" : logDrive.Mq,
	}
	for _, v := range dirveConfigArr {
		_, err := Call(dirveInterface, v,  m, _dir)
		if err != nil {
			fmt.Println("method invoke error:", err)
		}
	}

}

/**
 反射执行函数
 */
func Call(m map[string]interface{}, name string, params ...interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		return nil, errors.New("the number of input params not match!")
	}
	in := make([]reflect.Value, len(params))
	for k, v := range params {
		in[k] = reflect.ValueOf(v)
	}
	return f.Call(in), nil
}

