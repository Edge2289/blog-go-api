package model

import (
	"fmt"
	"sync"
)

type User struct {
	Model
	LID   int    `gorm:index:index_like_id`
	TTile string `gorm:"size:50,index:index_t_tile"` // 大小为 50
}

/**
创建用户数据表
*/
func CreatedUserTable() bool {
	if !Eloquent.HasTable(&User{}) {
		if err := Eloquent.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&User{}).Error; err != nil {
			panic(err)
		}
	}
	return true
}

/**
随机生成数据
*/
func RandUserData() int {
	var wg sync.WaitGroup

	for i := 0; i < 10; i = i + 1 {
		wg.Add(1)
		go func(n int) {

			defer wg.Add(-1)
			for y := 1; y <= 10; y++ {
				user := &User{
					LID:   y,
					TTile: "测试",
				}

				fmt.Print(i)
				if err := Eloquent.Create(user).Error; err != nil {
					fmt.Print(err)
				}
			}
		}(i)

	}

	wg.Wait()
	fmt.Println("END")
	return 1
}
