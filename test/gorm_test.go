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

	"Go-WebCreate/model"

	sqlConn "Go-WebCreate/utils/DB/mariadb"
)

func TestSQLConnect(t *testing.T) {
	if err := sqlConn.NewSqlConn("root", "00000000", "localhost", "3306", "first", "test"); err != nil {
		t.Error(err)
		fmt.Println("数据库连接失败")
	}
	fmt.Println("数据库连接成功")
}


func TestTableIsCreate(t *testing.T) {
	TestSQLConnect(t)
	sql := sqlConn.GetSqlConn()
	// 将User的IsCreate()方法返回值设置为false
	sqlConn.CreateTable(sql,&model.User{});
}