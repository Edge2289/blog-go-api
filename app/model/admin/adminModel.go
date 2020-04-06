package admin

import (
	"blog-go-api/app/model"
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

	RoleData []RoleModel `gprm:""`
	Users []User `gorm:"foreignkey:LID;AssociationForeignKey:ID" json:"users"`
}

// 设置表的名称
func (AdminModel) TableName() string {
	return "V_admin"
}

func (admin *AdminModel) GetAdminList(AdminModel, error) {
	var adminList AdminModel
	db.Limit(100000).Find(&rows)
}