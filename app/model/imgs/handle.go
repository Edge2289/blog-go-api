package imgs

import (
	db "blog-go-api/app/model"
	//timeUtils "blog-go-api/utils/time"
	"time"
)

/**
CREATE TABLE `V_imgs_categoryt` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL COMMENT '图片分类名称',
  `is_state` tinyint(3) DEFAULT NULL COMMENT '启用',
  `sort` int(3) DEFAULT NULL COMMENT '权重',
  `operator_id` int(10) DEFAULT NULL,
  `operator_name` varchar(30) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

type ImgsCategory struct {
	Id      int    `json:"id"`
	Name    string `gorm:"column:name;" json:"name"`
	IsState int    `gorm:"column:is_state;" json:"is_state"`
	Sort    int    `gorm:"column:sort;" json:"sort"`

	OperatorId   int    `gorm:"column:operator_id;"`
	OperatorName string `gorm:"column:operator_name;"`

	CreateTime time.Time  `json:"createTime" gorm:"column:created_at"`
	UpdateTime time.Time  `json:"updateTime" gorm:"column:updated_at"`
	DeleteTime *time.Time `json:"deleteTime" gorm:"column:deleted_at"`

	Imgs []Imgs
}

/**
  获取单个图片分类
*/
func (ImgsCate *ImgsCategory) GetImgsCategoryOne() (ImgsCategory, error) {

	var imgsCateData ImgsCategory

	err := db.Eloquent.Model(&imgsCateData).Where("is_state = ? and name = ?", 1, ImgsCate.Name).Order("sort desc").First(&imgsCateData).Error
	if err != nil {
		return imgsCateData, err
	}
	return imgsCateData, err
}

/**
  获取多个图片分类
*/
func (ImgsCate *ImgsCategory) GetImgsCategorys() ([]ImgsCategory, error) {

	var imgsCateData []ImgsCategory

	err := db.Eloquent.Model(&imgsCateData).Where("is_state = ?", 1).Order("sort desc").Find(&imgsCateData).Error
	if err != nil {
		return imgsCateData, err
	}
	return imgsCateData, err
}

/**
新增图片分类
*/
func (ImgCate *ImgsCategory) AddImsCategory() (bool, error) {

	ImgCate.IsState = 1
	if ImgCate.Sort == 0 {
		ImgCate.Sort = 50
	}
	err := db.Eloquent.Create(&ImgCate).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
更新图片分类
*/
func (ImgsCate *ImgsCategory) UpdateImsCategory() (bool, error) {

	err := db.Eloquent.Model(&ImgsCate).Where("id = ? ", ImgsCate.Id).Updates(&ImgsCate).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
删除图片分类
*/
func (ImgsCate *ImgsCategory) DelImsCategory() (bool, error) {

	err := db.Eloquent.Model(&ImgsCate).Where("id = ? ", ImgsCate.Id).Delete(&ImgsCate).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
