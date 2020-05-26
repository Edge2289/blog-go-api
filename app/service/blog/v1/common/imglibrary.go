package common

import (
	imgsModel "blog-go-api/app/model/imgs"
)

/**
上传图片
*/
func UploadImg() {

}

/**
删除图片
*/
func DelImgs() {

}


/**
获取图片分组
*/
func GetImgGroup() (imgs[] imgsModel.ImgsCategory, err error) {

	var img imgsModel.ImgsCategory

	imgs, err = img.GetImgsCategorys()
	if err != nil {
		return imgs, err
	}
	return imgs, nil
}

/**
删除图片分组
*/
func DelImgGroup(delImg imgsModel.ImgsCategory) (bool, error) {

	bool, err := delImg.DelImsCategory()
	if err != nil {
		return false, err
	}
	if bool {
		return true, nil
	}
	return false, nil
}

/**
更新图片分组
*/
func UpdateImgGroup(updateImgs imgsModel.ImgsCategory) (bool, error) {

	bool, err := updateImgs.UpdateImsCategory()
	if err != nil {
		return false, err
	}
	if bool {
		return true, nil
	}
	return false, nil
}

/**
新增图片分组
*/
func AddImgGroup(addImg imgsModel.ImgsCategory) (bool, error) {

	data, err := addImg.AddImsCategory()

	if err != nil {
		return false, err
	}
	if data {
		return true, nil
	}
	return false, nil
}