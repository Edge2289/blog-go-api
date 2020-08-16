package article

import (
	"blog-go-api/app/model/category"
	timeI "blog-go-api/utils/time"
	"blog-go-api/utils/tools"
	"fmt"
	"sync"
	"testing"
)

/**
	獲取列表
 */
func TestCategoryList(t *testing.T)  {
	var category category.Category
	list, count , _ := category.GetList(1, 30)
	fmt.Print(count)
	fmt.Print(list)
}

/**
	新增
 */
func TestAddCategory(t *testing.T)  {

	var workResultLock sync.WaitGroup
	for ii := 0; ii < 1000;  ii ++{
		workResultLock.Add(1)
		go forAdd()
		fmt.Print("協程 開啓")
	}
	//主线程等待
	workResultLock.Wait()
	fmt.Print("結束")
}
func forAdd()  {
	var category category.Category
	for i := 0; i < 50000; i++ {
		category.Name = "ce自行車自行車shi"+ tools.IntToString(i)
		category.IsState = 1
		category.IsHome = 0
		category.Note = "ceshi1 note" + tools.IntToString(i) + "啊擦撒十大大蘇打撒旦撒大蘇打實打實的阿斯頓撒旦"
		category.Sort = 50
		category.OperatorId = 1
		category.OperatorName = "測試0"
		category.CreateTime = timeI.GetDatabaseDate()
		state, err := category.AddCategory()
		fmt.Print(state)
		fmt.Print(err)
	}
}

func TestUpdateCategory(y *testing.T)  {

	var category category.Category

	category.Id = 6051087
	category.Name = "ce自行車自行車shiasdas"
	category.IsState = 0
	category.IsHome = 0
	category.Sort = 60
	category.Note = "ceshi1 note 啊擦撒十大大蘇打撒旦撒大蘇打實打實的阿斯頓撒旦"
	category.UpdateTime = timeI.GetDatabaseDate()
	state , _ := category.UpdateCategory()
	fmt.Print(state)
}

