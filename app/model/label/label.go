package label

import (
	"time"
	db "blog-go-api/app/model"
)

/**
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `label` varchar(30) DEFAULT NULL COMMENT '标签名称',
  `color` varchar(20) DEFAULT NULL COMMENT '颜色 #aaaaaa',
  `is_state` tinyint(3) DEFAULT NULL COMMENT '标签是否开启',
  `operator_id` int(10) DEFAULT NULL,
  `operator_name` varchar(30) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
 */
type Label struct {

	Id       int `json:"id"`
	Label        string `gorm:"column:label;" json:"label"`
	Color        string `gorm:"column:color;" json:"color"`
	IsState        int `gorm:"column:is_state;" json:"is_state"`

	OperatorId        int `gorm:"column:operator_id;"`
	OperatorName        string `gorm:"column:operator_name;"`

	CreateTime time.Time `json:"createTime" gorm:"column:created_at"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:updated_at"`
	DeleteTime *time.Time `json:"deleteTime" gorm:"column:deleted_at"`
}

/**
 新增
 */
func (label Label) AddLabel() (bool, error){

	err := db.Eloquent.Debug().Model(&label).Create(&label).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
	更新
 */
func (label Label) UpdateLabel() (bool, error){

	err := db.Eloquent.Debug().Model(&label).Where("id = ?", label.Id).Update(&label).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
	软删除
 */
func (label Label) DelLabel() (bool, error) {

	err := db.Eloquent.Model(&label).Where("id = ? ", label.Id).Delete(&label).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
	获取列表
 */
func (label Label) GetLabel(page, pageSize int) ([]Label, error) {

	var labelData[] Label
	labelModel := db.Eloquent.Model(&label)
	if label.Label != "" {
		labelModel = labelModel.Where("label like ?", "%"+label.Label+"%")
	}
	if label.IsState != -1 {
		labelModel = labelModel.Where("is_state = ?",label.IsState )
	}

	error :=labelModel.Offset((page-1) * pageSize).Limit(pageSize).Find(&labelData).Error
	if error != nil{
		return labelData, error
	}
	return labelData, error
}