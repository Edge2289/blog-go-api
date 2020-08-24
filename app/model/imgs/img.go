package imgs

import (
	db "blog-go-api/app/model"
	"time"
)

//CREATE TABLE `V_imgs` (
//`id` int(11) NOT NULL AUTO_INCREMENT,
//`url` varchar(255) DEFAULT NULL COMMENT '地址',
//`cate_id` int(11) DEFAULT NULL COMMENT '分类id',
//`domain` varchar(50) DEFAULT NULL COMMENT '域名',
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='图片库';

type Imgs struct {
	Id     int    `json:"id"`
	Url    string `gorm:"column:url;" json:"url"`
	CateId int    `gorm:"column:cate_id;" json:"cate_id"`
	Domain string `gorm:"column:domain;" json:"domain"`

	OperatorId   int    `gorm:"column:operator_id;"`
	OperatorName string `gorm:"column:operator_name;"`

	CreateTime time.Time  `json:"createTime" gorm:"column:created_at"`
	UpdateTime time.Time  `json:"updateTime" gorm:"column:updated_at"`
	DeleteTime *time.Time `json:"deleteTime" gorm:"column:deleted_at"`
}

/**
新增图片
*/
func (imgs *Imgs) AddImg() (bool, error) {

	err := db.Eloquent.Debug().Model(&imgs).Create(&imgs).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
删除如萍
*/
func (imgs *Imgs) DelImg() (bool, error) {

	err := db.Eloquent.Model(&imgs).Where("id = ? ", imgs.Id).Delete(&imgs).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
获取图片总数
*/

func (imgs *Imgs) GetImgsCount() (int, error) {
	var count int
	dbModel := db.Eloquent.Model(&count)
	if imgs.CateId != 0 {
		dbModel.Where("cate_id = ? ", imgs.CateId)
	}
	// 分页
	err := dbModel.Count(&count).Error
	if err != nil {
		return count, err
	}
	return count, err
}

/**
获取图片
*/
func (imgs *Imgs) GetImgs(page, pageSize int) ([]Imgs, error) {
	var imgsData []Imgs
	dbModel := db.Eloquent.Model(&imgs)
	if imgs.CateId != 0 {
		dbModel = dbModel.Where("cate_id = ? ", imgs.CateId)
	}
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 30
	}
	// 分页
	err := dbModel.Offset((page - 1) * pageSize).Limit(pageSize).Find(&imgsData).Error
	if err != nil {
		return imgsData, err
	}
	return imgsData, err
}

/**
更新图片分类
*/
func (imgs *Imgs) UpdateImsCategory() (bool, error) {

	err := db.Eloquent.Model(&imgs).Where("id = ? ", imgs.Id).Updates(&imgs).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
