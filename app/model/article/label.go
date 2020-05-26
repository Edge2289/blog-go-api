package article


import (
	"time"
	db "blog-go-api/app/model"
)

/**
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `aritcle_id` int(11) DEFAULT NULL COMMENT '文章id',
  `label_id` int(11) DEFAULT NULL COMMENT '标签id',
 */
type ArticleLabel struct {

	Id       int `json:"id"`
	ArticleId        int `gorm:"column:article_id;" json:"article_id"`
	LabelId        int `gorm:"column:label_id;" json:"label_id"`

	OperatorId        int `gorm:"column:operator_id;"`
	OperatorName        string `gorm:"column:operator_name;"`

	CreateTime time.Time `json:"createTime" gorm:"column:created_at"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:updated_at"`
	DeleteTime *time.Time `json:"deleteTime" gorm:"column:deleted_at"`
}