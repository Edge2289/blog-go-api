package router

import (
	"blog-go-api/app/handle/blog/V1/Admin"
)

func AdminRouter(base string)  {

	r := Routers.Group(apiRouter + "v1/" + base)

	r.POST("/admin", Admin.AddAdmin)
	//r.Use(middleware.CheckAuthRole())
	//{
		//r.GET("/deptList", GetDeptList)
	//}
}