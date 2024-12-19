/**
 * @Author tanchang
 * @Description 连接配置
 * @Date 2024/7/11 16:14
 * @File:  config
 * @Software: GoLand
 **/

package conf

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	"Go-WebCreate/utils/DB"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("读取配置文件环境失败" + err.Error())
	}

	//logrus配置
	logLevel, _ := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	logrus.SetLevel(logrus.Level(logLevel))
	logrus.SetReportCaller(true)

	//连接数据库
	// utils.DBUntil(os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_ADDR"), os.Getenv("DB_NAME"), os.Getenv("TABLE_NAME"))

	//连接redis
	db , _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	DB.NewRedisConn(os.Getenv("REDIS_ADDR"), os.Getenv("REDIS_PWD"),os.Getenv("REDIS_PORT"),db)
}
