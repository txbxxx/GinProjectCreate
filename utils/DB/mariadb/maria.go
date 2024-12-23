/**
 * @Author tanchang
 * @Description Mysql/Mariadb 连接工具类
 * @Date 2024/7/11 15:57
 * @File:  maria
 * @Software: GoLand
 **/

package sql

import (
	"Go-WebCreate/model"
	logg "Go-WebCreate/utils/log"
	"fmt"
	"log"
	"os"
	"reflect"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	sqlConn *gorm.DB
	dbOnce  sync.Once
)

type DbSql struct {
	DbPwd       string
	DbHost      string
	DbUser      string
	DbPort      string
	DbName      string
	TablePrefix string
}

func (t *DbSql) Connect() *gorm.DB {
	connAddr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", t.DbUser, t.DbPwd, t.DbHost, t.DbPort, t.DbName)

	// databases := t.DbUser + ":" + t.DbPwd + "@tcp(" + t.DbHost + ")/" + t.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	//配置数据库
	gormConfig := &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   t.TablePrefix,
			SingularTable: true,
		},
		Logger: newLog(),
	}
	
	db, err := gorm.Open(mysql.Open(connAddr), gormConfig)

	if err != nil {
		logg.Errorln("数据库连接错误: ", err.Error())
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
	dbOnce.Do(func ()  {
		sqlConn = NewDbSql(user, pwd, host, dbName, port, prefix).Connect()
	})
	sqlSet, err := sqlConn.DB()
	if err != nil {
		logg.Errorln("数据库连接错误: ", err.Error())
		return err
	}
	// 设置连接池
	sqlSet.SetMaxIdleConns(10)

	//  设置最大打开连接数
	sqlSet.SetMaxOpenConns(20)
	return nil
}

func GetSqlConn() *gorm.DB {
	if sqlConn == nil {
		logg.Errorln("请先创建数据库连接")
		return nil
	}
	logg.Infoln("获取数据库连接成功")
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
func CreateTable(Conn *gorm.DB, models ...model.CreateTable) error {
	if Conn == nil {
		logg.Errorln("数据库连接为空")
		return fmt.Errorf("数据库连接为空")
	}
	for _, model := range models {
		if model == nil {
			logg.Errorln("model为空")
			return fmt.Errorf("model为空")
		}

		if !model.IsCreate(){
			name := reflect.ValueOf(model).Elem().Type().Name()
			logg.Warningf("model: %s 该model不适合创建, 请传入正确的模型",  name)
			continue
		}

		logrus.Debugln("创建表: ", model)
		if err := Conn.AutoMigrate(model);err != nil {
			logrus.Errorln("创建表失败: ", err.Error())
			return err
		}
	}
	return nil
}
