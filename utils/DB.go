/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 15:57
 * @File:  DB
 * @Software: GoLand
 **/

package utils

import (
	"Go-WebCreate/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

var DB *gorm.DB

// DBUntil 用于连接数据库
func DBUntil(DBUser, DBPwd, DBAddr, DBName, TablePrefix string) {
	databases := DBUser + ":" + DBPwd + "@tcp(" + DBAddr + ")/" + DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	//配置数据库
	db, err := gorm.Open(mysql.Open(databases), &gorm.Config{
		SkipDefaultTransaction: false, //禁用事务
		NamingStrategy: schema.NamingStrategy{ //命名策略
			TablePrefix:   TablePrefix,
			SingularTable: true, //禁用复数名称
		},
	})

	sqlDB, err := db.DB()

	// 设置连接池
	sqlDB.SetMaxIdleConns(10)

	//  设置最大打开连接数
	sqlDB.SetMaxOpenConns(20)
	if err != nil {
		log.Println("数据库连接失败", err.Error())
	}

	DB = db
	//创建表
	CreateTable()

}

// CreateTable 使用自动迁移创建表
func CreateTable() {
	err := DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Println("创建表失败", err.Error())
	}
}
