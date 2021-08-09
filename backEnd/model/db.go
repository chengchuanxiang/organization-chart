package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"oyccx/utils"
	"time"
)

var Db *gorm.DB
var err error
func InitDb() {
	Db, err = gorm.Open(utils.Db, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	))
	if err != nil {
		fmt.Printf("连接数据库失败，请检查错误:",err)
	}
	//禁用默认表名的复数形式
	Db.SingularTable(true)
	//自动迁移
	Db.AutoMigrate(&Data{})
	//设置连接池中的最大闲置链接数
	Db.DB().SetMaxIdleConns(10)

	//设置数据库的最大连接量
	Db.DB().SetMaxOpenConns(100)

	//设置连接的最大可复用时间
	Db.DB().SetConnMaxLifetime(10 * time.Second)

	//db.Close()

}