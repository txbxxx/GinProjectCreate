/**
 * @Author tanchang
 * @Description Redis 连接工具类
 * @Date 2024/7/11 16:31
 * @File:  Redis
 * @Software: GoLand
 **/

package Redis

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var redisConn *redis.Client

type RedisNoSql struct {
	RedisPW        string
	RedisHost      string
	RedisDefaultDB int
	RedisProt      string
}

func (t *RedisNoSql) Connect() *redis.Client {
	addr := fmt.Sprintf("%s:%s", t.RedisHost, t.RedisProt)
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: t.RedisPW,
		DB:       t.RedisDefaultDB,
	})
}

func NewRedisNoSql(RDBHost, RDBPwd, RDBPORT string, RDBDefaultDB int) *RedisNoSql {

	return &RedisNoSql{
		RedisHost:      RDBHost,
		RedisProt:      RDBPORT,
		RedisPW:        RDBPwd,
		RedisDefaultDB: RDBDefaultDB,
	}
}

// 创建redis连接
func NewRedisConn(RDBHost, RDBPwd, RDBPORT string, RDBDefaultDB int)  error{
	redisNoSql := NewRedisNoSql(RDBHost, RDBPwd, RDBPORT, RDBDefaultDB)
	redisConn = redisNoSql.Connect()
	if _,err:= redisConn.Ping(context.Background()).Result();err != nil {
		logrus.Error("redis连接失败: ",err.Error())
		CloseRedisConn()
		return err
	}
	return nil
}

// 获取redis连接
func GetRedisConn() *redis.Client {
	if redisConn != nil {
		logrus.Info("获取redis连接成功")
		return redisConn
	}
	logrus.Error("获取redis连接失败,请先调用NewRedisConn方法创建redis连接")
	return nil
}

// 关闭redis连接
func CloseRedisConn() {
	redisConn.Close()
}
