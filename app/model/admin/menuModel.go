package admin

import (
	db "blog-go-api/app/model"
	"blog-go-api/utils/time"
)

/**
	菜单的字段
    菜单的无限极分类数据结构
*/
type Menu struct {
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

	CreateTime string `json:"createTime" gorm:"column:created_at"`
	UpdateTime string `json:"updateTime" gorm:"column:updated_at"`
	DeleteTime string `json:"deleteTime" gorm:"column:deleted_at"`
	Children     []Menu `json:"Children"`
}

/**
 菜单角色关联表
 */
type MenuRoleModel struct {
	RoleId string `gorm:"column:role_id;" json:"role_id"`
	MenuId string `gorm:"column:menu_id;" json:"menu_id"`
}


/********************     CURD     **************************/
/**
	获取菜单列表
	menuSearch 搜索条件
 */
func (menuSearch Menu) GetMenuList() ([]Menu, error) {
	var menuList[] Menu
	table := db.Eloquent.Table("V_menu")
	if menuSearch.Label != "" {
		table.Where("label = ?", menuSearch.Label)
	}
	err := table.Find(&menuList).Error
	if err != nil {
		return nil, err
	}
	//menuChildren := genMenuTree(menuList,0)
	return menuList, nil
}

/**
 新增菜单
 */
func (menuData Menu) InsertMenuList() (bool, error) {

	menuData.CreateTime = time.GetCurrentDate()
	err := db.Eloquent.Table("V_menu").Create(&menuData).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
 更新菜单数据
 */
func (menuData Menu) UpdateMenuList() (bool, error) {
	menuData.UpdateTime = time.GetCurrentDate()
	err := db.Eloquent.Table("V_menu").Where("id = ? ", menuData.Id).Updates(&menuData).Error
	if err != nil {
		return false, err
	}
	return true, nil;
}

/**
 删除菜单结构
 */
func (menuData Menu) DeleteMenuList() (bool, error) {

	menuData.DeleteTime = time.GetCurrentDate()
	err := db.Eloquent.Table("V_menu").Where("id = ? ", menuData.Id).Updates(&menuData).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

/********************     CURD     **************************/


/**
 根据过来的用户id 获取
 */
func GetAdminMenu(uids int) ([]Menu, error) {
	var menuData []Menu

	var roleMenuModel []MenuRoleModel
	db.Eloquent.Raw("select menu_id from V_role_menu where role_id in (select role_id from V_admin_role where u_id = ?)", uids).Find(&roleMenuModel)
	var roleMenuData []string
	for _,v := range roleMenuModel {
		roleMenuData = append(roleMenuData, v.MenuId)
	}
	db.Eloquent.Where("id in (?)", roleMenuData).Find(&menuData)
	menuChildren := genMenuTree(menuData,0)
	return menuChildren, nil
}

/**
 供给请求方 jwt 鉴权使用
 应当做一下缓存 redis
 */
func GetAdminJwtMenu(uids int) ([]Menu, error) {
	var menuData []Menu

	var roleMenuModel []MenuRoleModel
	db.Eloquent.Raw("select menu_id from V_role_menu where role_id in (select role_id from V_admin_role where u_id = ?)", uids).Find(&roleMenuModel)
	var roleMenuData []string
	for _,v := range roleMenuModel {
		roleMenuData = append(roleMenuData, v.MenuId)
	}
	db.Eloquent.Where("id in (?)", roleMenuData).Find(&menuData)
	menuChildren := genMenuTree(menuData,0)
	return menuChildren, nil
}

/**
 将菜单生成 无限极分类
 */

func genMenuTree(menueData []Menu, pid int) []Menu{

	var menuChile []Menu
	for _, v := range menueData {
		if pid == v.Pid {
			v.Children = genMenuTree(menueData, v.Id)
			menuChile = append(menuChile, v)
		}
	}
	return menuChile
}