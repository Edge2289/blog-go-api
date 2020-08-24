package time

import (
	"strconv"
	"time"
)

// 获取当前的时间 - 字符串
func GetDatabaseDate() time.Time {
	return time.Now()
}

// 获取当前的时间 - 字符串
func GetCurrentDate() string {
	return time.Now().Format("2006/01/02 15:12:12")
}

// 获取当前小时 - 字符串
func GetDateHours() string {
	return strconv.Itoa(time.Now().Hour()) //小时
}

// 获取当前年月日 - 字符串
func GetDateDMY() string {
	return time.Now().Format("20060102")
}
