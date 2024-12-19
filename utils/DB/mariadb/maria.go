/**
 * @Author tanchang
 * @Description Mysql/Mariadb 连接工具类
 * @Date 2024/7/11 15:57
 * @File:  maria
 * @Software: GoLand
 **/

package sql

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var sqlConn *gorm.DB

type DbSql struct {
	DbPwd       string
	DbHost      string
	DbUser      string
	DbPort      string
	DbName      string
	TablePrefix string
}

func (t *DbSql) Connect() *gorm.DB {
	//定义gorm的日志配置
	newLogger := newLog()

	connAddr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", t.DbUser, t.DbPwd, t.DbHost, t.DbPort, t.DbName)

	// databases := t.DbUser + ":" + t.DbPwd + "@tcp(" + t.DbHost + ")/" + t.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	//配置数据库
	// TODO 定制化
	db, err := gorm.Open(mysql.Open(connAddr), &gorm.Config{
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

func NewSqlConn(user, pwd, host, port, dbName, prefix string) error {
	sqlConn = NewDbSql(user, pwd, host, dbName, port, prefix).Connect()
	sqlSet, err := sqlConn.DB()
	if err != nil {
		logrus.Error("数据库连接错误: ", err.Error())
		return err
	}
	// 设置连接池
	sqlSet.SetMaxIdleConns(10)

	//  设置最大打开连接数
	sqlSet.SetMaxOpenConns(20)
	return nil
}

func GetSqlConn() *gorm.DB {
	return sqlConn
}

// 创建gorm日志设置
// TODO 定制化
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
func CreateTable(Conn *gorm.DB, models ...interface{}) error {
	if Conn == nil {
		logrus.Errorln("数据库连接为空")
		return fmt.Errorf("数据库连接为空")
	}
	for _, model := range models {
		if model == nil {
			logrus.Errorln("model为空")
			return fmt.Errorf("model为空")
		}
		logrus.Debugln("创建表: ", model)
		if err := Conn.AutoMigrate(model);err != nil {
			logrus.Errorln("创建表失败: ", err.Error())
			return err
		}
	}
	return nil
}
