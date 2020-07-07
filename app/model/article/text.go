package article

import (
	"time"
)
/**
`id` int(11) NOT NULL AUTO_INCREMENT,
  `article_id` int(11) DEFAULT NULL COMMENT '文章id',
  `text` text COMMENT '文章内容',
  `markdown` text COMMENT '文章的markdown',
 */
type ArticleText struct {

	Id       int `json:"id"`
	ArticleId        int `gorm:"column:article_id;" json:"article_id"`
	Text        string `gorm:"column:text;" json:"text"`
	Markdown        string `gorm:"column:markdown;" json:"markdown"`

	OperatorId        int `gorm:"column:operator_id;"`
	OperatorName        string `gorm:"column:operator_name;"`

	CreateTime time.Time `json:"createTime" gorm:"column:created_at"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:updated_at"`
	DeleteTime *time.Time `json:"deleteTime" gorm:"column:deleted_at"`
}

/**
	新增
 */
func AddArticleText(articleText ArticleText) (bool, error) {

	err := db.Model(&articleText).Create(&articleText).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func UpdateArticleText(articleText ArticleText) (bool, error) {

	err := db.Where("id = ?", articleText.Id).Update(&articleText).Error
	if err != nil {
		return false, err
	}
	return true, nil
}