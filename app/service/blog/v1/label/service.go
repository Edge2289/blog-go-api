package label

import (
	labelModel "blog-go-api/app/model/label"
)

/**
新增
*/
func AddLabel(label labelModel.Label) (bool, error) {

	i, error := label.AddLabel()
	return i, error
}

/**
更新
*/
func UpdateLabel(label labelModel.Label) (bool, error) {

	i, error := label.UpdateLabel()
	return i, error
}

/**
删除
*/
func DelLabel(label labelModel.Label) (bool, error) {

	i, error := label.DelLabel()
	return i, error
}

/**
查询
*/
func GetLabel(label labelModel.Label, page, pageSize int) ([]labelModel.Label, int, error) {
	data, count, err := label.GetLabel(page, pageSize)
	return data, count, err
}
