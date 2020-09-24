package qiniu

import (
	"blog-go-api/app/config"
	timeGet "blog-go-api/utils/time"
	"strconv"
	"time"

	"context"
	"fmt"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

var (
	accessKey = config.AccessKey // 七牛的accessKey 去七牛后台获取
	secretKey = config.SecretKey // 七牛的secretKey 去七牛后台获取
	bucket    = config.BUCKET    // 上传空间 去七牛后台创建
)

/**
上传七牛云的图片
*/
func Upload(imgData []byte) (string, error) {
	// 鉴权
	mac := qbox.NewMac(accessKey, secretKey)

	// 上传策略
	putPolicy := storage.PutPolicy{
		Scope:   bucket,
		Expires: 7200,
	}

	// 获取上传token
	upToken := putPolicy.UploadToken(mac)

	// 上传Config对象
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan //指定上传的区域
	cfg.UseHTTPS = false           // 是否使用https域名
	cfg.UseCdnDomains = false      //是否使用CDN上传加速

	// 七牛key
	newImgName := "blog/" + timeGet.GetDateDMY() + "/" + strconv.Itoa(int(time.Now().Unix())) + ".png"

	// 构建上传的对象
	base64Uploader := storage.NewBase64Uploader(&cfg)
	fmt.Println(base64Uploader)
	ret := storage.PutRet{}

	//图片base64格式的数据 注意 需要去掉 前面类似data:image/png;base64,的数据

	imgData = imgData[23:]
	fmt.Println(imgData)
	err := base64Uploader.Put(context.Background(), &ret, upToken, newImgName, imgData, nil)
	if err != nil {
		fmt.Println("上传文件失败,原因:", err)
		return newImgName, err
	}
	fmt.Println("上传成功,key为:", ret.Key)
	return newImgName, nil
}

/**
删除七牛云的图片
*/
func DelImgs(imgName string) (bool, error) {

	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	cfg.Zone = &storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)
	// 删除七牛云的图片
	err := bucketManager.Delete(bucket, imgName)
	if err != nil {
		return false, err
	}
	return true, nil
}
