package system

/**

  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `label` varchar(60) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称（暂时没有用处）',
  `pid` int(11) NOT NULL DEFAULT '0' COMMENT '父类id',
  `icon` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '图标',
  `api_path` varchar(100) NOT NULL COMMENT 'Api请求的路径\n',
  `method` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '请求的方式  post',
  `vue_router_path` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '页面路径',
  `vue_router_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '页面的名字',
  `vue_router_component` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci DEFAULT NULL COMMENT '页面地址 默认不加@',
  `is_function` tinyint(3) unsigned DEFAULT '0' COMMENT '是否功能健 0 不是 1 是',
  `is_used` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '启用 0 不启用 1启用（暂时没有用处）',
  `operator_id` int(10) unsigned DEFAULT '0' COMMENT '操作员id',
  `operator_name` varchar(30) DEFAULT NULL COMMENT '操作员名称',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
 菜单模型
*/
type Menu struct {

	Id    int `gorm:"column:id;primary_key"`
	Label        string `gorm:"column:label;type:varchar(60);not null;"`
	Pid        int `gorm:"column:pid;type:int(11);not null;"`
	Icon        string `gorm:"column:icon;type:varchar(32);not null;"`
	ApiPath        string `gorm:"column:api_path;default:0;"`
	Method        string `gorm:"column:method;type:varchar(11);"`
	VueRouterPath        string `gorm:"column:vue_router_path;default:1;"`
	VueRouterName        string `gorm:"column:vue_router_name;"`
	VueRouterComponent   string `gorm:"column:vue_router_component;"`
	IsFunction        int `gorm:"column:is_function;"`
	IsUsed        int `gorm:"column:is_used;"`

}
