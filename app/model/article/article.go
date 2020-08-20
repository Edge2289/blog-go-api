package article

import (
	dbModel "blog-go-api/app/model"
	"fmt"
	"time"
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
	Id           int    `json:"id"`
	Title        string `gorm:"column:title;" json:"title"`
	Describe     string `gorm:"column:describe;" json:"describe"`
	Img          string `gorm:"column:img;" json:"img"`
	Nick         string `gorm:"column:nick;" json:"nick"`
	CateId       int    `gorm:"column:cate_id;" json:"cate_id"`
	Introduction string `gorm:"column:introduction;" json:"introduction"`
	IsComment    int    `gorm:"column:is_comment;" json:"is_comment"`
	IsState      int    `gorm:"column:is_state;" json:"is_state"`
	ClickNum     int    `gorm:"column:click_num;" json:"click_num"`
	ReadNum      int    `gorm:"column:read_num;" json:"read_num"`

	OperatorId   int    `gorm:"column:operator_id;"`
	OperatorName string `gorm:"column:operator_name;"`

	CreateTime time.Time  `json:"createTime" gorm:"column:created_at"`
	UpdateTime time.Time  `json:"updateTime" gorm:"column:updated_at"`
	DeleteTime *time.Time `json:"deleteTime" gorm:"column:deleted_at"`

	/**
	关联列表
	*/
	// 标签
	LabelData []ArticleLabel `gorm:"foreignkey:ArticleId;AssociationForeignKey:ID" json:"label_data"`
	// 点击量
	ClickList []ArticleClick `gorm:"foreignkey:ArticleId;AssociationForeignKey:ID" json:"click_list"`
	// 文章明细
	TextData ArticleText `gorm:"foreignkey:ArticleId;AssociationForeignKey:ID" json:"text_data"`
	// 评论列表
	CommentsList ArticleComments `gorm:"foreignkey:ArticleId;AssociationForeignKey:ID" json:"comments_list"`
}

var db = dbModel.Eloquent

/**
获取查询搜索的列表
可搜索参数：文章分类，文章类别，文章名称
按照创建时间排序
*/
func (article Article) GetSearchList(label string, page, pageSize int) ([]Article, int, error) {

	var articleList []Article
	articleModel := db.Model(&articleList)
	if article.CateId != 0 {
		// 分类搜索条件不为0
		articleModel = articleModel.Where("cate_id = ? ", article.CateId)
	}
	if label != "" {
		var articleLabelList []ArticleLabel
		err := db.Model(&articleLabelList).Debug().Select("article_id").Where("label_id = ?", label).Find(&articleLabelList).Error
		if err != nil {
			return articleList, 0, err
		}
		var ids []interface{}
		for _, v := range articleLabelList {
			ids = append(ids, v.ArticleId)
		}
		articleModel = articleModel.Where("id in (?) ", ids)
	}
	if article.Title != "" {
		// 标题搜索条件不为0
		articleModel = articleModel.Where("title like ?", "%"+article.Title+"%")
	}
	if article.CateId != 0 {
		// 标题搜索条件不为0
		articleModel = articleModel.Where("cate_id = ?", article.CateId)
	}
	if article.Nick != "" {
		// 标题搜索条件不为0
		articleModel = articleModel.Where("nick like ? ", "%"+article.Nick+"%")
	}
	if article.IsState != 0 {
		// 标题搜索条件不为0
		articleModel = articleModel.Where("is_state = ?", article.IsState)
	}

	err := articleModel.Preload("LabelData").Offset((page - 1) * pageSize).Limit(pageSize).Find(&articleList).Error
	if err != nil {
		return articleList, 0, err
	}
	var count int
	articleModel.Count(&count)
	// 需要连表查询
	return articleList, count, nil
}

/**
获取文章的明细
	过来的是文章id
*/
func (article Article) GetArticleDetail() (articleDetail Article, err error) {

	err = db.Model(&article).Preload("LabelData").Preload("ClickList").Preload("TextData").Preload("CommentsList").Where("id = ? and is_state = 1", article.Id).Find(&articleDetail).Error
	return articleDetail, err
}

/**
新增文章
使用事务操作
*/
func (article Article) AddArticleDetail() (bool, error) {

	fmt.Println("-----------")
	fmt.Println(article.Id)
	err := db.Model(&article).Create(&article).Error
	if err != nil {
		return false, err
	}
	fmt.Println("-----------")
	fmt.Println(article.Id)

	// 添加文章明细
	//article.TextData.ArticleId = article.Id
	//_, err = AddArticleText(article.TextData)
	//if err != nil {
	//	return false, err
	//}
	//// 添加文章标签
	//c, err := AddArticleLabel(article.LabelData)
	//if !b || !c {
	//	tx.Rollback()
	//	return false, err
	//}

	//// 发生错误时回滚事务
	//tx.Rollback()

	fmt.Println(11)
	return true, nil
}

/**
修改文章
*/

/**
删除文章
*/
