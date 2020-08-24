package adminValidators

import (
	"blog-go-api/app/model/admin"
)

/**
登陆的验证器
*/
func LoginValidators(loginData admin.LoginData) string {
	// 用户名
	if loginData.Username == "" || !(len(loginData.Username) > 0 && len(loginData.Username) < 20) {
		return "-10004"
	}
	// 密码
	if loginData.Password == "" || !(len(loginData.Password) > 0 && len(loginData.Password) < 20) {
		return "-10005"
	}
	// 验证码
	if loginData.Code == "" || !(len(loginData.Code) > 0 && len(loginData.Code) < 5) {
		return "-10006"
	}
	return ""
}
