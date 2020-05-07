package main

import (
	"blog-go-api/app/config"
	imgsModel "blog-go-api/app/model/imgs"
	//"blog-go-api/router"
	"fmt"
)

func main() {
	fmt.Print("------------------------------------ ")
	fmt.Print("---------   WEB BLOG API   --------- ")
	fmt.Print("---                              --- ")
	fmt.Print("---  appName:"+config.AppName+"  --- ")
	fmt.Print("---  port:  "+config.AppPort+"   --- ")
	fmt.Print("------------------------------------ ")
	fmt.Print("------------------------------------ ")

	var img imgsModel.ImgsCategory
	var imgs imgsModel.ImgsCategory
	imgs.Id = 2
	imgs.Name = "测试分类"
	data, err := img.GetImgsCategorys()
	datas, err := imgs.UpdateImsCategory()
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println("data", data)
	fmt.Println("datas", datas)
	//router.Run()
}

