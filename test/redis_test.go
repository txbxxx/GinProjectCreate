/**
 * @Author tanchang
 * @Description redis测试
 * @Date 2024/7/11 22:59
 * @File:  redisTest
 * @Software: GoLand
 **/

package test

import (
	"context"
	"fmt"
	"testing"
	"time"

	redisConn "Go-WebCreate/utils/DB/redis"
)

func TestRedisConnect(t *testing.T) {
	if err := redisConn.NewRedisConn("127.0.0.1", "", "6379", 0); err != nil {
		t.Error("redis连接错误")
	}
	conn := redisConn.GetRedisConn()
	if _, err := conn.Ping(context.Background()).Result(); err == nil {
		fmt.Println("redis 连接成功")
	}
}


func TestRedisSet(t *testing.T) {
	ctx := context.Background()
	TestRedisConnect(t)
	conn := redisConn.GetRedisConn()
	_,err := conn.Set(ctx,"name","txbxxx",time.Minute).Result()
	if err != nil {
		t.Error("redis set错误")
	}
	s := conn.Get(ctx,"name").Val()
	fmt.Println(s)
}
