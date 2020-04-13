package model

import (
	"fmt"
	"blog-go-api/app/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var Eloquent *gorm.DB
// 基本模型的定义
type Model struct {
	Id        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)


	dbType = config.DbType
	dbName = config.DbDatabase
	user = config.DbUserName
	password = config.DbPassword
	host = config.DbHost+":"+config.DbHost
	tablePrefix = ""

	Eloquent, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(Eloquent *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	Eloquent.SingularTable(true)
	Eloquent.DB().SetMaxIdleConns(10)
	Eloquent.DB().SetMaxOpenConns(100)
}


func CloseDB() {
	defer Eloquent.Close()
}


//func init(){
//	db, err := gorm.Open("mysql", "root:123456@/sys?charset=utf8&parseTime=True&loc=Local")
//	if err != nil {
//		fmt.Print(err)
//	}
//}



// create 创建表
//if !db.HasTable(&Like{}) {
//	if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(&Like{}).Error; err != nil {
//		panic(err)
//	}
//}
//var wg sync.WaitGroup
//
//for i := 0; i < 100; i = i + 1 {
//
//	wg.Add(1)
//
//	go func(n int) {
//
//		defer wg.Add(-1)
//
//		for i := 1; i <= 10000; i++ {
//			ip := "127.0.0.1"
//			ua := "11.111.11"
//			title := "测试"
//
//			like := &Like{
//				Ip:        ip,
//				Ua:        ua,
//				Title:     title,
//				Hash:      1,
//				CreatedAt: time.Now(),
//			}
//
//			fmt.Print(i)
//			if err := db.Create(like).Error; err != nil {
//				fmt.Print(err)
//			}
//		}
//
//	}(i)
//
//}
//
//wg.Wait()
//fmt.Println("END")
// 插入
//for i := 1; i <= 10000; i++ {
//	ip := "127.0.0.1"
//	ua := "11.111.11"
//	title := "测试"
//
//	like := &Like{
//		Ip:        ip,
//		Ua:        ua,
//		Title:     title,
//		Hash:      1,
//		CreatedAt: time.Now(),
//	}
//
//	fmt.Print(i)
//	if err := db.Create(like).Error; err != nil {
//		fmt.Print(err)
//	}
//}




//like := &Like{}
// 查询
//var rows [] Like
//data := db.Limit(100000).Find(&rows)
//data.Scan(like)
//fmt.Println(like)
//fmt.Println(rows)


//ret := res.Scan(&Like{})
//if res.RecordNotFound() {
//	fmt.Print(12312)
//} else {
//	fmt.Print(res.Rows())
//	fmt.Print(res.Scan(&Like{}))
//}
//fmt.Print(res.Rows())



//user := User{}
//fmt.Print(user)
//res := db.First(&user)
//fmt.Print("%+v", res)
//defer db.Close()