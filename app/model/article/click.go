package article

import (
	"time"
)

/**
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `u_id` int(11) DEFAULT NULL COMMENT '用户id',
  `article_id` int(11) DEFAULT NULL COMMENT '文章id',
*/

type ArticleClick struct {
	Id        int `json:"id"`
	UId       int `gorm:"column:u_id;" json:"u_id"`
	ArticleId int `gorm:"column:article_id;" json:"article_id"`

	OperatorId   int    `gorm:"column:operator_id;"`
	OperatorName string `gorm:"column:operator_name;"`

	CreateTime time.Time  `json:"createTime" gorm:"column:created_at"`
	UpdateTime time.Time  `json:"updateTime" gorm:"column:updated_at"`
	DeleteTime *time.Time `json:"deleteTime" gorm:"column:deleted_at"`
}
