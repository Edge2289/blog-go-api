package admin

import (
	"blog-go-api/app/config"
	db "blog-go-api/app/model"
	"blog-go-api/common"
	"fmt"
)

/**
登陆的前台请求参数
binding 代表着参数需要验证
*/
type LoginData struct {
	Username string `form:"user_name" json:"user_name" binding:"required"`
	Password string `form:"pwd" json:"pwd" binding:"required"`
	Code string `form:"code" json:"code" binding:"required"`
	UUID string `form:"uuid" json:"uuid" binding:"required"`
}

/**
 后台管理员模型
 */
type Admin struct {

	Id    int `gorm:"column:id;primary_key"`
	LoginName        string `gorm:"column:login_name;type:varchar(30);not null;index:admin_user_index"`
	Name        string `gorm:"column:name;type:varchar(30);not null;"`
	Password        string `gorm:"column:password;type:varchar(32);not null;"`
	Sex        int `gorm:"column:sex;default:0;"`
	Phone        string `gorm:"column:phone;type:varchar(11);"`
	IsUsed        string `gorm:"column:is_used;default:1;"`
	DepartmentId        int `gorm:"column:password;"`
	OperatorId        int `gorm:"column:operator_id;"`
	OperatorName        string `gorm:"column:operator_name;"`

	RoleData []AdminRole `gprm:"foreignkey:UId;AssociationForeignKey:Id" json:"role_data"`
	MenuData []MenuModel `gorm:"foreignkey:LID;AssociationForeignKey:ID" json:"menu_data"`
	Token string
}

/**
 登陆模块
 */
func (loginData *LoginData) LoginGetUserList() (Admin, error) {
	var adminData Admin

	pwdenty:= EncryptionPwd(loginData.Password)

	fmt.Println(loginData.Username, pwdenty)

	err := db.Eloquent.Model(&adminData).Where("login_name = ? and password = ?", loginData.Username, pwdenty).First(&adminData).Error
	if err != nil {
		return adminData, err
	}

	// 清空密码，避免前台返回数据显示
	adminData.Password = ""
	adminData.RoleData, _ = GetRoleData(adminData.Id)
	adminData.MenuData, _ = GetAdminMenu(adminData.Id)
	// 查询数据库是否有这个东西
	return adminData, nil
}

func (admin *Admin) Logout(token string) (status bool, err error) {
	return
}

/**
 获取管理云的全部账号
 filter  包含了搜索条件以及查询的页数
 */
func (admin *Admin) GetAdminList(filter map[string]interface{}) (status bool, err error) {
	//var adminList AdminModel
	//db.Eloquent.Limit(100000).Find(&rows)
	return false, nil
}

/**
 新增管理员
 */
func (admin *Admin) AddAdminList() (bool, error) {

	admin.Password = EncryptionPwd(admin.Password) // 密码加密
	admin.OperatorId = 1
	admin.OperatorName = "测试添加"

	if err := db.Eloquent.Create(admin).Error; err != nil {
		return false, err
	}
	return false, nil
}

/**
 编辑管理员数据
 */
func (admin *Admin) EditAdminList (id int64) (update Admin, err error) {
	return
}

/**
 删除管理员 （软删除）
 */
func (admin *Admin) DelAdminList (id int64) (success bool, err error) {
	return false, nil
}

/**
 加密管理云密码
	md5(md5($pass).$salt);
 */
func EncryptionPwd(pwd string) (p string) {

	// 密匙
	salt := config.AppSalt
	return common.MD5(common.MD5(pwd)+salt)
}