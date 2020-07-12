package article


import (
	"time"
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

/**
新增标签id
*/
func AddArticleLabel(articleLabel []ArticleLabel) (bool, error) {

	// 循环插入标签
	for _, v := range articleLabel  {
		err := db.Model(&v).Create(&v).Error
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

/**
	更新标签
 */

func UpdateArticleLabel(articleLabel []ArticleLabel, articleId int) (bool, error) {

	// 首先查询出之前的，对比
	//tx.Query("update V_article_label set deleted_at = ? where article_id = ?", time.Time{}, articleId)
	// 循环插入标签
	//b, err := AddArticleLabel(articleLabel)
	return false, nil
}