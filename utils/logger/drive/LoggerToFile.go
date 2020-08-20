package drive

import (
	"blog-go-api/utils/time"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

//调用os.MkdirAll递归创建文件夹
func createFile(filePath string) error {
	if !isExist(filePath) {
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

/**
 * 新建一个日志文件
 * 输入一个信息  以及文件名
 */
func createAllFile(fileName string) error {
	fileNameArr := strings.Split(fileName, "/")
	var fileStringName string = ""
	for _, value := range fileNameArr {
		if strings.Index(value, ".") != -1 {
			fileStringName = fileStringName + value
			if !isExist(fileStringName) {
				_, err := os.Create(fileStringName)
				return err
			}
		} else {
			fileStringName = fileStringName + value + "/"
			_err := createFile(fileStringName)
			if _err != nil {
				return _err
			}
		}
	}
	return nil
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

/**
保存一个文件日志
*/
func File(m map[string]interface{}, _dir string) {

	_dir = _dir + "/" + time.GetDateDMY() + "/request-" + time.GetDateHours() + ".log"
	// 新建文件 或者是 目录
	createAllFile(_dir)
	// 写入文件 或者 新建文件夹
	src, err := os.OpenFile(_dir, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Print("err", err)
		return
	}

	logger := logrus.New()
	//设置输出
	logger.Out = src

	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{})

	// 日志格式
	logger.Infoln(m)
}
