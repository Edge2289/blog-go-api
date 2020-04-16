package admin

import (
	db "blog-go-api/app/model"
)

/**
	菜单的字段
    菜单的无限极分类数据结构
*/
type MenuModel struct {
	Id       int             `json:"id"`
	Label        string `gorm:"column:label;" json:"label"`
	Pid        int `gorm:"column:pid;" json:"pid"`
	Icon        string `gorm:"column:icon;" json:"icon"`
	ApiPath        string `gorm:"column:api_path;" json:"api_path"`
	Method        string `gorm:"column:method;" json:"method"`
	VueRouterPath        string `gorm:"column:vue_router_path;" json:"vue_router_path"`
	VueRouterName        string `gorm:"column:vue_router_name;" json:"vue_router_name"`
	VueRouterComponent        string `gorm:"column:vue_router_component;" json:"vue_router_component"`
	IsFunction        int8 `gorm:"column:is_function;" json:"is_function"`
	IsUsed        int8 `gorm:"column:is_used;" json:"is_used"`
	Children     []MenuModel `json:"Children"`
}

/**
 菜单角色关联表
 */
type MenuRoleModel struct {
	RoleId string `gorm:"column:role_id;" json:"role_id"`
	MenuId string `gorm:"column:menu_id;" json:"menu_id"`
}

/**
 根据过来的用户id 获取
 */
func GetAdminMenu(uids int) ([]MenuModel, error) {

	var menueData []MenuModel
	db.Eloquent.Raw("select * from V_menu where id in (select menu_id from V_role_menu where role_id in (select role_id from V_admin_role where u_id = ?)) and is_used = 1 and deleted_at is NULL", uids).Find(&menueData)
	menuChildren := genMenuTree(menueData,0)
	return menuChildren, nil
}

/**
 供给请求方 jwt 鉴权使用
 应当做一下缓存 redis
 */
func GetAdminJwtMenu(uids int) ([]MenuModel, error) {
	var menuData []MenuModel
	db.Eloquent.Raw("select id, method, api_path from V_menu where id in (select menu_id from V_role_menu where role_id in (select role_id from V_admin_role where u_id = 1)) and is_used = ? and deleted_at is NULL and api_path <> \"\" and method <> \"\"", uids).Find(&menuData)
	return menuData, nil
}

/**
 将菜单生成 无限极分类
 */

func genMenuTree(menueData []MenuModel, pid int) []MenuModel{

	var menuChile []MenuModel
	for _, v := range menueData {
		if pid == v.Pid {
			v.Children = genMenuTree(menueData, v.Id)
			menuChile = append(menuChile, v)
		}
	}
	return menuChile
}