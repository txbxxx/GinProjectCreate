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

	sqlConn "Go-WebCreate/utils/DB/mariadb"
)

func TestRedis(t *testing.T) {
	if err := sqlConn.NewSqlConn("root", "00000000", "localhost", "3306", "first", "test"); err != nil {
		t.Error(err)
		fmt.Println("数据库连接失败")
	}
	fmt.Println("数据库连接成功")
}
