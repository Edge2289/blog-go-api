package admin

import (
	db "blog-go-api/app/model"
)

/**
`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `name` varchar(50) NOT NULL COMMENT '角色名称',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '状态 0 关闭 1 开始',
  `content` varchar(200) DEFAULT NULL COMMENT '角色描述',
  `operator_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '操作员id',
  `operator_name` varchar(30) DEFAULT NULL COMMENT '操作员名称',
*/
type Role struct {
	Id           int    `gorm:"column:id;"`
	Name         string `gorm:"column:name;"`
	Status       int32  `gorm:"column:status;"`
	Content      string `gorm:"column:content;"`
	OperatorId   int    `gorm:"column:operator_id;"`
	OperatorName string `gorm:"column:operator_name;"`
}

type AdminRole struct {
	Id           int    `gorm:"column:id;"`
	UId          int    `gorm:"column:u_id;"`
	RId          int    `gorm:"column:role_id;"`
	OperatorId   int    `gorm:"column:operator_id;"`
	OperatorName string `gorm:"column:operator_name;"`
}

/**
  获取用户的角色
*/
func GetRoleData(uid int) ([]AdminRole, error) {
	var roleData []AdminRole
	err := db.Eloquent.Model(&roleData).Where("u_id = ?", uid).Select("role_id").Find(&roleData).Error
	if err != nil {
		return roleData, err
	}
	return roleData, nil
}
