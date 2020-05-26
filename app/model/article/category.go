package article

import (
	"time"
	db "blog-go-api/app/model"
)
/**
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `title` varchar(50) DEFAULT NULL COMMENT '标题',
  `describe` varchar(100) DEFAULT NULL COMMENT '描述',
  `img` varchar(255) DEFAULT NULL COMMENT '配图',
  `nick` varchar(50) DEFAULT NULL COMMENT '作者昵称',
  `cate_id` int(11) DEFAULT NULL COMMENT '分类id',
  `introduction` varchar(100) DEFAULT NULL COMMENT '简介',
  `is_comment` tinyint(3) DEFAULT '1' COMMENT '是否开启评论',
  `is_state` tinyint(3) DEFAULT '1' COMMENT '是否发布',
  `click_name` int(11) DEFAULT '0' COMMENT '点赞数',
  `read_num` int(11) DEFAULT '0' COMMENT '阅读数',
*/
type Article struct {

	Id       int `json:"id"`
	Title        string `gorm:"column:title;" json:"title"`
	Describe     string `gorm:"column:describe;" json:"describe"`
	Img        string `gorm:"column:img;" json:"img"`
	Nick        string `gorm:"column:nick;" json:"nick"`
	CateId        int `gorm:"column:cate_id;" json:"cate_id"`
	Introduction        string `gorm:"column:introduction;" json:"introduction"`
	IsComment        int `gorm:"column:is_comment;" json:"is_comment"`
	IsState        int `gorm:"column:IsState;" json:"IsState"`
	ClickName        int `gorm:"column:click_name;" json:"click_name"`
	ReadNum        int `gorm:"column:read_num;" json:"read_num"`

	OperatorId        int `gorm:"column:operator_id;"`
	OperatorName        string `gorm:"column:operator_name;"`

	CreateTime time.Time `json:"createTime" gorm:"column:created_at"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:updated_at"`
	DeleteTime *time.Time `json:"deleteTime" gorm:"column:deleted_at"`
}
