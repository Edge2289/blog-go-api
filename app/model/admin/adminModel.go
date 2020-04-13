package admin

import (
	"blog-go-api/app/config"
	db "blog-go-api/app/model"
	"blog-go-api/common"
	"fmt"
)

/**
 后台管理员模型
 */
type AdminModel struct {

	LoginName        string `gorm:"column:login_name;type:varchar(30);not null;index:admin_user_index"`
	Name        string `gorm:"column:name;type:varchar(30);not null;"`
	Password        string `gorm:"column:password;type:varchar(32);not null;"`
	Sex        int `gorm:"column:sex;default:0;"`
	Phone        string `gorm:"column:phone;type:varchar(11);"`
	IsUsed        string `gorm:"column:is_used;default:1;"`
	DepartmentId        int `gorm:"column:password;"`
	OperatorId        int `gorm:"column:operator_id;"`
	OperatorName        string `gorm:"column:operator_name;"`

	//RoleData []RoleModel `gprm:""`
	//Role []RoleModel `gorm:"foreignkey:LID;AssociationForeignKey:ID" json:"role"`
	//Menu []MenuModel `gorm:"foreignkey:LID;AssociationForeignKey:ID" json:"users"`
}

// 设置表的名称
func (AdminModel) TableName() string {
	return "V_admin"
}

/**
 登陆模块
 */
func (admin *AdminModel) Login() ([] AdminModel, error) {
	fmt.Print("Login Handle")
	return nil, nil
}

func (admin *AdminModel) Logout(token string) (status bool, err error) {
	return
}

/**
 获取管理云的全部账号
 filter  包含了搜索条件以及查询的页数
 */
func (admin *AdminModel) GetAdminList(filter map[string]interface{}) (status bool, err error) {
	//var adminList AdminModel
	//db.Eloquent.Limit(100000).Find(&rows)
	return false, nil
}

/**
 新增管理员
 */
func (admin *AdminModel) AddAdminList() (bool, error) {

	admin.Password = EncryptionPwd(admin.Password) // 密码加密

	if err := db.Eloquent.Create(admin).Error; err != nil {
		return false, err
	}
	return false, nil
}

/**
 编辑管理员数据
 */
func (admin *AdminModel) EditAdminList (id int64) (update AdminModel, err error) {
	return
}

/**
 删除管理员 （软删除）
 */
func (admin *AdminModel) DelAdminList (id int64) (success bool, err error) {
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