/**
 * @Author tanchang
 * @Description //TODO
 * @Date 2024/7/11 22:59
 * @File:  gormtest
 * @Software: GoLand
 **/

package test

import (
	"fmt"
	"testing"

	"Go-WebCreate/utils"
	"Go-WebCreate/utils/DB"
)

func TestCreateUser(t *testing.T) {
	md5 := utils.GetMd5("123456")
	println(md5)
}

func TestRedis(t *testing.T) {
	if err := DB.NewRedisConn("127.0.0.1", "", "6379", 0); err != nil {
		t.Error(err)
		fmt.Println("连接redis失败")
	}
	fmt.Println("连接redis成功")
}
