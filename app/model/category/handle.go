package category

import (
	dbModel "blog-go-api/app/model"
	"time"
)
var db = dbModel.Eloquent

// 定義分類的結構體
type Category struct {
	Id        int `gorm:"column:id;" json:"id"`
	Name        string `gorm:"column:name;" json:"name"`
	Note        string `gorm:"column:note;" json:"note"`
	IsState        int `gorm:"column:is_state;" json:"is_state"`
	Sort 		int`gorm:"column:sort;" json:"sort"`
	IsHome 		int`gorm:"column:is_home;" json:"is_home"`

	OperatorId        int `gorm:"column:operator_id;"`
	OperatorName        string `gorm:"column:operator_name;"`

	CreateTime time.Time `json:"createTime" gorm:"column:created_at"`
	UpdateTime time.Time `json:"updateTime" gorm:"column:updated_at"`
	DeleteTime *time.Time `json:"deleteTime" gorm:"column:deleted_at"`
}

/**
  獲取列表
 */
func (category Category) GetList(page int, pageSize int) ([]Category, int, error) {
	var categoryList []Category
	categoryModel := db.Model(&categoryList)

	if category.Id != 0 {
		categoryModel = categoryModel.Where("id = ?", category.Id)
	}
	if category.Name != "" {
		categoryModel = categoryModel.Where("name like ?", category.Name+"%")
	}
	if category.IsState != -1 {
		categoryModel = categoryModel.Where("is_state = ?", category.IsState)
	}
	// 搜索條件
	err := categoryModel.Offset((page - 1) * pageSize).Limit(pageSize).Find(&categoryList).Error
	if err != nil {
		return categoryList, 0, err
	}
	var count int
	// 計算總數
	categoryModel.Count(&count)
	return categoryList, count, nil
}

/**
	新增
 */
func (category Category) AddCategory() (bool, error) {

	err := db.Model(&category).Create(&category).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

/**
更新
*/
func (category Category) UpdateCategory() (bool, error) {

	data := make(map[string]interface{})

	data["name"] = category.Name
	data["sort"] = category.Sort
	data["is_home"] = category.IsHome
	data["is_state"] = category.IsState
	data["note"] = category.Note

	err := db.Model(Category{}).Where("id = ?", category.Id).Updates(data).Error
	if err != nil {
		return false, err
	}
	return true, nil
}


/**
刪除
*/
func (category Category) DelCategory() (bool, error) {

	err := db.Model(&category).Delete(&category).Error
	if err != nil {
		return false, err
	}
	return true, nil
}