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

	"Go-WebCreate/model"

	"github.com/joho/godotenv"

	sqlConn "Go-WebCreate/utils/DB/mariadb"
	redisConn "Go-WebCreate/utils/DB/redis"
	"Go-WebCreate/utils/log"
)

func Init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("读取配置文件环境失败" + err.Error())
	}

	//初始化日志
	log.NewProjectLog()

	//连接数据库
	// utils.DBUntil(os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_ADDR"), os.Getenv("DB_NAME"), os.Getenv("TABLE_NAME"))
	sqlConn.NewSqlConn(os.Getenv("DB_USER"), os.Getenv("DB_PWD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"), os.Getenv("TABLE_NAME"))

	//连接redis
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	redisConn.NewRedisConn(os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PWD"), os.Getenv("REDIS_PORT"), db)

	// 初始化数据库
	conn := sqlConn.GetSqlConn()
	sqlConn.CreateTable(conn, &model.User{})
}
