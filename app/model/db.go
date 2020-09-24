package model

import (
	"blog-go-api/app/config"
	"fmt"
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
	fmt.Println("数据库连接开始")

	dbType = config.DbType
	dbName = config.DbDatabase
	user = config.DbUserName
	password = config.DbPassword
	host = config.DbHost
	tablePrefix = "V_"

	Eloquent, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))

	if err != nil {
		log.Println("数据库连接失败")
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(Eloquent *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	// 启用Logger，显示详细日志
	Eloquent.LogMode(true)
	Eloquent.SingularTable(true)
	Eloquent.DB().SetMaxIdleConns(10)
	Eloquent.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer Eloquent.Close()
}
