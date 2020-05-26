package article

import (
	"time"
	db "blog-go-api/app/model"
)
/**
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `u_id` int(11) DEFAULT NULL COMMENT '用户id',
  `comment` varchar(255) DEFAULT NULL COMMENT '评论内容',
  `is_state` tinyint(3) DEFAULT NULL COMMENT '是否开启显示',
  `p_id` int(3) DEFAULT NULL COMMENT '层级结构',
  `click` int(11) DEFAULT NULL COMMENT '点赞数',
 */

type ArticleComments struct {

	Id       int `json:"id"`
	UId        int `gorm:"column:u_id;" json:"u_id"`
	Comment    string `gorm:"column:comment;" json:"comment"`
	IsState    int `gorm:"column:is_state;" json:"is_state"`
	PId    int `gorm:"column:p_id;" json:"p_id"`
	Click    int `gorm:"column:click;" json:"click"`

	OperatorId        int `gorm:"column:operator_id;"`
	OperatorName        string `gorm:"column:operator_name;"`

	CreateTime time.Time `json:"createTime" gorm:"column:created_at"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:updated_at"`
	DeleteTime *time.Time `json:"deleteTime" gorm:"column:deleted_at"`
}