/**
 * @Author tanchang
 * @Description Mysql/Mariadb 连接工具类
 * @Date 2024/7/11 15:57
 * @File:  Db
 * @Software: GoLand
 **/

package DB

import (
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"Go-WebCreate/model"
)

var Db *gorm.DB

type DbSql struct {
	DbPwd       string
	DbHost      string
	DbUser      string
	DbPort      string
	DbName      string
	TablePrefix string
}

func NewDbSql(DbUser, DbPwd, DbHost, DbName, DbPort, TablePrefix string) *DbSql {
	return &DbSql{
		DbPwd:       DbPwd,
		DbUser:      DbUser,
		DbPort:      DbPort,
		DbName:      DbName,
		DbHost:      DbHost,
		TablePrefix: TablePrefix,
	}
}

func (t *DbSql) Connect() *gorm.DB {
	//定义gorm的日志配置
	newLogger := newLog()

	databases := t.DbUser + ":" + t.DbPwd + "@tcp(" + t.DbHost + ")/" + t.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	//配置数据库
	db, err := gorm.Open(mysql.Open(databases), &gorm.Config{
		SkipDefaultTransaction: false, //禁用事务
		NamingStrategy: schema.NamingStrategy{ //命名策略
			TablePrefix:   t.TablePrefix,
			SingularTable: true, //禁用复数名称
		},
		Logger: newLogger,
	})

	if err != nil {
		logrus.Error("数据库连接错误: ", err.Error())
		return nil
	}
	return db
}

// DbUntil 用于连接数据库
func DbUntil(DbUser, DbPwd, DbAddr, DbName, TablePrefix string) *gorm.DB {
	//定义gorm的日志配置
	newLogger := newLog()

	databases := DbUser + ":" + DbPwd + "@tcp(" + DbAddr + ")/" + DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	//配置数据库
	db, err := gorm.Open(mysql.Open(databases), &gorm.Config{
		SkipDefaultTransaction: false, //禁用事务
		NamingStrategy: schema.NamingStrategy{ //命名策略
			TablePrefix:   TablePrefix,
			SingularTable: true, //禁用复数名称
		},
		Logger: newLogger,
	})

	if err != nil {
		logrus.Error("数据库连接错误: ", err.Error())
	}

	sqlDb, err := db.DB()

	// 设置连接池
	sqlDb.SetMaxIdleConns(10)

	//  设置最大打开连接数
	sqlDb.SetMaxOpenConns(20)
	if err != nil {
		logrus.Println("数据库连接失败: ", err.Error())
	}

	Db = db
	//创建表
	CreateTable()

}

func newLog() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			ParameterizedQueries:      true,
			Colorful:                  false,
		},
	)
	return newLogger
}

// CreateTable 使用自动迁移创建表
func CreateTable() {
	err := Db.AutoMigrate(&model.User{})
	if err != nil {
		logrus.Println("创建表失败", err.Error())
		return
	}
}
