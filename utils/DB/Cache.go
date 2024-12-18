/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 16:31
 * @File:  Cache
 * @Software: GoLand
 **/

package utils

import (
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var Cache *redis.Client

// RedisUtils redis连接
func RedisUtils(RDBAddr, RDBPwd, RDBDefaultDB string) {
	// 将字符串转换成int
	RDB, err := strconv.Atoi(RDBDefaultDB)
	if err != nil {
		fmt.Println("转换地址失败")
	}

	//连接redis
	Cache = redis.NewClient(&redis.Options{
		Addr:     RDBAddr,
		Password: RDBPwd,
		DB:       RDB,
	})
}
